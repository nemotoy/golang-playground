package roundtrip

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_sampleRoundtrip(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		w.WriteHeader(200)
		w.Write([]byte("success!"))
	}))
	defer ts.Close()

	c := &http.Client{Transport: &LogTransport{count: 100}}
	resp, err := c.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "text/plain" {
		t.Fatalf("expected text/plain, got %d", resp.StatusCode)
	}

	if v, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fatal(err)
	} else if string(v) != "success!" {
		t.Fatalf("expected %q, got %q", "success!", v)
	}
}

type testTransport struct{}

func (t *testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer([]byte("test")))}, nil
}

func Test_testRoundtrip(t *testing.T) {

	c := &http.Client{Transport: &testTransport{}}
	req, err := http.NewRequest(http.MethodGet, "test", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if v, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fatal(err)
	} else if string(v) != "test" {
		t.Fatalf("expected %q, got %q", "test", v)
	}
}
