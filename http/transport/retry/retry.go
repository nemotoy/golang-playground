package roundtrip

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	// copy the request body because of rewind it
	var b []byte
	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		b = body
	}
	for {
		counter++
		res, err := t.transport().RoundTrip(req)
		if err != nil || (res != nil && res.StatusCode >= 500) {
			if counter == t.MaxRetries {
				return nil, fmt.Errorf("achieved given max retries, then failed to request: error = %v, stauts code = %d", err, res.StatusCode)
			}
			// wait a few seconds
			if b != nil {
				// need to assign a variable that satisfies the io.ReadCloser interface
				req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}
			continue
		}
		return res, nil
	}
}
