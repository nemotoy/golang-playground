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
	RetryFunc  func(res *http.Response, err error) bool
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.RetryFunc == nil {
		t.RetryFunc = defaultRetry
	}
	var (
		counter int = -1
		// copy the request body because of rewind it
		b   []byte
		err error
	)
	b, err = getReqBody(req)
	if err != nil {
		return nil, err
	}
	for {
		counter++
		res, err := t.transport().RoundTrip(req)
		if t.RetryFunc(res, err) {
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

func defaultRetry(res *http.Response, err error) bool {
	return err != nil || (res != nil && shouldRetryStatus(res.StatusCode))
}

func shouldRetryStatus(status int) bool {
	return status == http.StatusInternalServerError || status == http.StatusBadGateway
}

func shouldRetryWithSwitch(status int) bool {
	switch status {
	case http.StatusInternalServerError, http.StatusBadGateway:
		return true
	}
	return false
}

func getReqBody(req *http.Request) ([]byte, error) {
	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, nil
}
