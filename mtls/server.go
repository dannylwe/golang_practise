package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func main() {
	http.HandleFunc("/hello", helloHandler)

	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// create TLS config and enable client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:         ":9000",
		TLSConfig:    tlsConfig,
		ReadTimeout:  6 * time.Second,
		WriteTimeout: 6 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("starting server on port 9000")
	// listen to HTTPS connections
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}
