package chain

import (
	"net/http"
	"sync"
)

func ClassifyTransport(t http.RoundTripper) http.RoundTripper {
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
	// log.Printf("#FirstTransport.flag: %v\n", t.f.f)
	return ClassifyTransport(t.Transport).RoundTrip(req)
}

type SecondTransport struct {
	Transport http.RoundTripper
	f         *flag
}

func (t *SecondTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.f.mu.Lock()
	t.f.f = true
	defer t.f.mu.Unlock()
	// log.Printf("#SecondTransport.flag: %v\n", t.f.f)
	return ClassifyTransport(t.Transport).RoundTrip(req)
}

// TODO: fix implementation to improve a versatility
func NewChainedTransports(f *flag) *http.Client {
	return &http.Client{
		Transport: &FirstTransport{
			Transport: &SecondTransport{f: f},
			f:         f,
		},
	}
}

type hkey string
type hval string

func (k hkey) String() string { return string(k) }
func (k hval) String() string { return string(k) }

type headTransport struct {
	Transport http.RoundTripper
	h         map[hkey]hval
}

func (t *headTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.h {
		req.Header.Add(k.String(), v.String())
	}
	// log.Printf("#headTransport.header: %v\n", req.Header)
	return ClassifyTransport(t.Transport).RoundTrip(req)
}

type headTransportWithF struct {
	Transport http.RoundTripper
	h         map[hkey]hval
	f         *flag
}

func (t *headTransportWithF) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.h {
		req.Header.Add(k.String(), v.String())
	}
	t.f.mu.Lock()
	t.f.f = false
	t.f.mu.Unlock()
	return ClassifyTransport(t.Transport).RoundTrip(req)
}
