package roundtrip

import (
	"fmt"
	"net/http"
)

func (t *Transport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

type Transport struct {
	Transport  http.RoundTripper
	MaxRetries int // if above it, stop executing a request
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	var counter int = -1
	for {
		counter++
		res, err := t.transport().RoundTrip(req)
		if err != nil || (res != nil && res.StatusCode >= 500) {
			if counter == t.MaxRetries {
				return nil, fmt.Errorf("achieved given max retries, then failed to request: error = %v, stauts code = %d", err, res.StatusCode)
			}
			// wait a few seconds
			continue
		}
		return res, nil
	}
}
