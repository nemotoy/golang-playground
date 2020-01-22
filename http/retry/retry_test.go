package retry

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"context"

	"fmt"

	"github.com/cenkalti/backoff"
	"github.com/google/go-cmp/cmp"
)

type response struct {
	S string
	N int
}

var (
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := dummyResponse()
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Fprintf(os.Stderr, "failed to encode: %#v", err)
		}
	})
	dummyResponse = func() response {
		return response{"hello", 10}
	}
)

func Test_Retryable(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	c := ts.Client()
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	var r *http.Response
	ctx := context.Background()
	err = backoff.Retry(func() (err error) {
		r, err = c.Do(req)
		if err == nil && r.StatusCode == http.StatusOK {
			return nil
		}
		r.Body.Close()
		return err
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
	if err != nil {
		t.Fatal(err)
	}
	var res response
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		t.Fatal(err)
	}
	got := res
	want := dummyResponse()
	if !cmp.Equal(want, got) {
		t.Errorf("got = %v; want %v", got, want)
	}
}
