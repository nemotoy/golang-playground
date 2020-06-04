package main

import (
	"io"
	"log"
	"net/http"
)

// refer to https://golang.org/pkg/net/http/#ListenAndServeTLS
// generate signed .pem files
// go run $GOROOT/src/crypto/tls/generate_cert.go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Received a request: ", req)
		_, _ = io.WriteString(w, "Hello, TLS!\n")
	})

	if err := http.ListenAndServeTLS(":8443", "./cert.pem", "./key.pem", nil); err != nil {
		log.Fatal(err)
	}
}
