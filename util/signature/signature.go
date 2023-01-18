package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/spf13/pflag"
)

var (
	flags                           *pflag.FlagSet
	pkgFile, privateKey, passphrase string
)

func main() {
	crypto.GetTime()

	// Read the command line flags.
	flags.Parse(os.Args[1:])
	if flags.NFlag() == 0 {
		flags.Usage()
		os.Exit(2)
	}
	pkgFile = flags.Lookup("package").Value.String()
	privateKey = flags.Lookup("private-key").Value.String()
	passphrase = flags.Lookup("passphrase").Value.String()
	if pkgFile == "" || privateKey == "" || passphrase == "" {
		log.Fatal("package, private-key and passphrase must all be set")
	}

	// Get the chart.yaml file and its contents.
	f, err := os.Open(pkgFile)
	if err != nil {
		log.Fatal("could not open package: %w", err)
	}
	defer f.Close()

	val, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("could not read package: %w", err)
	}
	raw := bytes.NewBuffer(val)

	chartYAML, err := extractChartYAML(raw)
	if err != nil {
		log.Fatal("could not extract Chart.yaml from package: ", err)
	}

	//Calculate the SHA256 sum of the (zipped) package.
	h := sha256.New()
	h.Write(val)
	checksum := fmt.Sprintf("%x", h.Sum(nil))

	// Generate the provenance file.
	pkgParts := strings.Split(pkgFile, "/")
	p := bytes.NewBuffer(nil)
	p.Write(chartYAML)
	p.Write([]byte("\n...\nfiles:\n"))
	p.Write([]byte(fmt.Sprintf("  %s: sha256:%s", pkgParts[len(pkgParts)-1], checksum)))
	prov := strings.ReplaceAll(p.String(), "- ", " - - ") // Replace starting dashes as per RFC 4880 (https://www.rfc-editor.org/rfc/rfc4880#section-7.1)

	// Sign the provenance file and write to the prov file.
	privKey, err := os.ReadFile(privateKey)
	if err != nil {
		log.Fatal("could not read private key: %w", err)
	}
	armored, err := helper.SignCleartextMessageArmored(string(privKey), []byte(passphrase), prov)
	fmt.Println(string(privKey))
	fmt.Println(passphrase)
	fmt.Println(prov)
	if err != nil {
		log.Fatal("could not sign provenance file: %w", err)
	}
	err = os.WriteFile(fmt.Sprintf("%s.prov", pkgFile), []byte(armored), 0644)
	if err != nil {
		log.Fatal("could not write provenance file: %w", err)
	}
}

func init() {
	flags = pflag.NewFlagSet("signature", pflag.ExitOnError)

	flags.String("package", "", "Packaged Helm chart (.tgz)")
	flags.String("private-key", "", "Locked GPG private key file")
	flags.String("passphrase", "", "Passphrase for the private key")

	flags.Usage = func() {
		_, err := os.Stderr.WriteString(
			fmt.Sprintf("Usage: helm-sign [OPTIONS]\n%s", flags.FlagUsages()))
		if err != nil {
			panic(err)
		}
	}
}

func extractChartYAML(r io.Reader) ([]byte, error) {
	var chartYAML []byte
	val, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("could not read package: %w", err)
	}
	raw := bytes.NewBuffer(val)
	// Decompress the archive.
	gz, err := gzip.NewReader(raw)
	if err != nil {
		return nil, err
	}
	v, err := ioutil.ReadAll(gz)
	if err != nil {
		return nil, fmt.Errorf("could not decompress archive: %w", err)
	}
	gz.Close()
	tr := tar.NewReader(bytes.NewBuffer(v))
	for {
		header, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if strings.HasSuffix(header.Name, "Chart.yaml") {
			chartYAML, err = io.ReadAll(tr)
			if err != nil {
				return nil, fmt.Errorf("could not read Chart.yaml: %w", err)
			}
			break
		}
	}
	if len(chartYAML) == 0 {
		return nil, fmt.Errorf("no or empty Chart.yaml found in package")
	}
	return chartYAML, nil
}
