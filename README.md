# http-dev-server

A simple HTTP development server written in Golang.

### Prerequisites
This repo can be used in many ways.

1. Run the Golang code directly; requires Go v1.18 and above.
2. Run using Docker; requires Docker Engine.
3. Run using Helm: requires Helm v3.8.0 and above.
  - Optional: If you want to sign and publish your own packages, then you need the [Helm GPG plugin](https://github.com/technosophos/helm-gpg).

## Usage

Run directly from docker.

```bash
$ docker run --rm -p 8080:8080 krishnaiyer/http-dev-server
```

## Helm

1. Setup a Kubernetes Cluster.

2. Install the Traefik CRDs.

```bash
$ kubectl apply -f https://raw.githubusercontent.com/traefik/traefik/v2.9/docs/content/reference/dynamic-configuration/kubernetes-crd-definition-v1.yml
```

3. Pull the Helm Charts

```bash
$ helm pull oci://registry-1.docker.io/krishnaiyer/http-dev-server-helm --version <version>
```

### Signing Helm Packages

Install [Helm GPG plugin](https://github.com/technosophos/helm-gpg).

List the secret keys in your keyring.

```bash
$ gpg --list-secret-keys
```

Sign by indicating the key and the package name.

```bash
$ KEY=ABCDEF... OCI_TAG=<tag> make helm.sign
```

## License

The contents of this repository are provided `as-is` under the terms of the [Apache 2.0](./LICENSE) license.
