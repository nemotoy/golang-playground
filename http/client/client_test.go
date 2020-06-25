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

/*
	## When is the request body closed?
	* request URL is nil
	https://github.com/golang/go/blob/master/src/net/http/client.go#L596
	* call the uerr and a reqBodyClosed is false. (ex, calls c.send() previously)
	https://github.com/golang/go/blob/master/src/net/http/client.go#L617
	* c.send() may be close req.Body
	https://github.com/golang/go/blob/master/src/net/http/client.go#L719
		* send() close req.Body when parameters are nil
		https://github.com/golang/go/blob/334752dc8207d6d19d9fb1a99d2e97f7d326c82a/src/net/http/client.go#L204
	* finish client.do()
	https://github.com/golang/go/blob/master/src/net/http/client.go#L737

*/
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
