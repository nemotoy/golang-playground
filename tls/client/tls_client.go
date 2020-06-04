package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var secureSkipped = flag.Bool("s", false, "")

// nolint: gosec
func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCert, err := ioutil.ReadFile("./cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal(err)
	}

	c := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: *secureSkipped,
			RootCAs:            certPool,
			Certificates:       []tls.Certificate{cert},
		},
	}}
	resp, err := c.Get("https://localhost:8443/")
	if err != nil {
		log.Printf("failed request: %#v", err)
		if e, ok := err.(*url.Error); ok {
			log.Fatalf("in detail: %+v", e.Unwrap())
		}
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response body: ", string(body))
}
