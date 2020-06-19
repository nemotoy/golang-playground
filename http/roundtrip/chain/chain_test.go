package chain

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func Test_ChainRoundtrip(t *testing.T) {
	f := &flag{mu: &sync.Mutex{}}
	c := NewChainedTransports(f)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !f.f {
			t.Fatal("f is false, expects true")
		}
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(200)
		w.Write([]byte("success!"))
	}))
	defer ts.Close()

	resp, err := c.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if v, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fatal(err)
	} else if string(v) != "success!" {
		t.Fatalf("expected %q, got %q", "success!", v)
	}
}

/*
	$ go test -count 2 -benchmem -bench .
	goos: darwin
	goarch: amd64
	pkg: github.com/nemotoy/golang-playground/http/roundtrip/chain
	Benchmark_headTransport-4                   2740            511275 ns/op           26907 B/op        132 allocs/op
	Benchmark_headTransport-4                   1137           1241009 ns/op           25685 B/op        130 allocs/op
	Benchmark_headTransportNonChain-4            592           1774655 ns/op           25013 B/op        130 allocs/op
	Benchmark_headTransportNonChain-4            549           1826221 ns/op           25245 B/op        131 allocs/op
	PASS
	ok      github.com/nemotoy/golang-playground/http/roundtrip/chain       5.704s
*/
func Benchmark_headTransport(b *testing.B) {

	f := &flag{mu: &sync.Mutex{}}
	hm := map[hkey]hval{
		"keyA": "valA",
	}
	c := &http.Client{
		Transport: &headTransport{
			Transport: &FirstTransport{f: f},
			h:         hm,
		},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("success!"))
	}))
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := c.Get(ts.URL)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_headTransportNonChain(b *testing.B) {

	f := &flag{mu: &sync.Mutex{}}
	hm := map[hkey]hval{
		"keyA": "valA",
	}
	c := &http.Client{
		Transport: &headTransportWithF{
			h: hm,
			f: f,
		},
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("success!"))
	}))
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := c.Get(ts.URL)
		if err != nil {
			b.Fatal(err)
		}
	}
}
