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
