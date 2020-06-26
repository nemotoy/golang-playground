package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ss struct {
	S string
}

func Test_Roundtrip(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(200)
		w.Write([]byte("hello"))
	}))
	defer ts.Close()
	var v ss
	b, err := json.Marshal(&v)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, ts.URL, bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}

	c := &http.Client{Timeout: 1 * time.Millisecond}
	resp, err := c.Do(req)
	if err != nil {
		// t.Fatalf("#Do; req: %v. error: %#v", req, err)
		return
	}
	defer resp.Body.Close()

	if v, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Fatal(err)
	} else if string(v) != "hello" {
		t.Fatalf("expected %q, got %q", "hello", v)
	}
}
