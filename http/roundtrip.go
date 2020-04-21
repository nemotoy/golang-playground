package roundtrip

import (
	"log"
	"net/http"
)

func (t *LogTransport) transport() http.RoundTripper {
	if t.Transport == nil {
		return http.DefaultTransport
	}
	return t.Transport
}

type LogTransport struct {
	Transport http.RoundTripper
}

func (t *LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("log: %v", req)
	return t.transport().RoundTrip(req)
}
