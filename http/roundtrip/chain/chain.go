package chain

import (
	"log"
	"net/http"
	"sync"
)

func transport(t http.RoundTripper) http.RoundTripper {
	if t == nil {
		return http.DefaultTransport
	}
	return t
}

type flag struct {
	mu *sync.Mutex
	f  bool
}

type FirstTransport struct {
	Transport http.RoundTripper
	f         *flag
}

func (t *FirstTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.f.mu.Lock()
	t.f.f = false
	t.f.mu.Unlock()
	log.Printf("#FirstTransport.flag: %v\n", t.f.f)
	return transport(t.Transport).RoundTrip(req)
}

type SecondTransport struct {
	Transport http.RoundTripper
	f         *flag
}

func (t *SecondTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.f.mu.Lock()
	defer t.f.mu.Unlock()
	t.f.f = true
	log.Printf("#SecondTransport.flag: %v\n", t.f.f)
	return transport(t.Transport).RoundTrip(req)
}
