package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	address := os.Getenv("PORT")
	if address == "" {
		address = ":8080"
	}
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Method: %v\n", r.Method)
		fmt.Printf("Headers: %v\n", r.Header)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}
		defer r.Body.Close()
		if len(body) > 0 {
			fmt.Printf("Body: %v\n", string(body))
		}
		w.WriteHeader(http.StatusAccepted)
	})
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	s := &http.Server{
		Addr:           address,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Start server: %v", address)
	log.Fatal(s.ListenAndServe())
}
