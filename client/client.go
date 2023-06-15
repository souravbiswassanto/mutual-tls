package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	cert, err := ioutil.ReadFile("/home/user/go/src/github.com/souravbiswassanto/mutual-tls/cert/ca.crt")
	if err != nil {
		log.Fatalf("could not open certificate file: %v\n", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)
	certificate, err := tls.LoadX509KeyPair("/home/user/go/src/github.com/souravbiswassanto/mutual-tls/cert/client.crt", "/home/user/go/src/github.com/souravbiswassanto/mutual-tls/cert/client.key")
	if err != nil {
		log.Fatalf("could not load certificate: %v\n", err)
	}
	client := http.Client{
		Timeout: time.Minute * 3,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{certificate},
			},
		},
	}

	resp, err := client.Get("https://localhost:9090")
	if err != nil {
		log.Fatalf("error making get request %v\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response: %v\n", err)
	}
	fmt.Println(string(body))
}
