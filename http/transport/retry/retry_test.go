package roundtrip

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Roundtrip(t *testing.T) {
	counter := 0
	wantRetries := 3
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counter++
		if counter == wantRetries {
			w.WriteHeader(200)
			w.Write([]byte("success"))
		} else {
			w.WriteHeader(500)
			w.Write([]byte("failed"))
		}
	}))
	defer ts.Close()

	c := &http.Client{Transport: &Transport{MaxRetries: wantRetries}}
	resp, err := c.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if wantRetries != counter {
		t.Fatalf("expected %d, got %d", wantRetries, counter)
	}
	if v, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fatal(err)
	} else if string(v) != "success" {
		t.Fatalf("expected %q, got %q", "success", v)
	}
}
