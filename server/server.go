package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("New Request -- home directory ")
		fmt.Fprintf(writer, "Hello World\n")
		//fmt.Println(writer, "hello world")
	})
	caCertFile, err := ioutil.ReadFile("/home/user/go/src/github.com/souravbiswassanto/mutual-tls/cert/ca.crt")
	if err != nil {
		log.Fatalf("error reading CA certificate: %v", err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCertFile)

	server := http.Server{
		Addr:    ":9090",
		Handler: handler,
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  certPool,
			MinVersion: tls.VersionTLS12,
		},
	}
	if err := server.ListenAndServeTLS("/home/user/go/src/github.com/souravbiswassanto/mutual-tls/cert/server.crt", "/home/user/go/src/github.com/souravbiswassanto/mutual-tls/cert/server.key"); err != nil {
		log.Fatalf("error listening to port: %v\n", err)
	}
}
