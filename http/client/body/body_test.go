package crbody

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

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

func TestClientDoAfterClosedRequestBody(t *testing.T) {
	var v = struct {
		S string
	}{
		S: "ss",
	}
	b, err := json.Marshal(&v)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("request URL is nil", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}
		prvBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Request body: %s", prvBody)
		c := &http.Client{}
		resp, err := c.Do(req)
		if err == nil {
			t.Error(err)
		}
		if resp != nil {
			resp.Body.Close()
			t.Errorf("expects response is nil")
		}
		// Although execute closing the request body, requst body is not nil.
		// What processing does The Close()?
		if req.Body != nil {
			nxtBody, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}
			t.Errorf("expects request body is nil; type: %#v, body: %s", req.Body, nxtBody)
		}
	})
}
