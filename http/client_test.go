package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"context"

	"fmt"

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

func Test_ShouldRetryRequest(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	tsc := ts.Client()
	c := New(tsc)
	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()

	resp, err := c.ShouldRetryRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		t.Fatal(err)
	}
	got := res
	want := dummyResponse()
	if !cmp.Equal(want, got) {
		t.Errorf("got = %v; want %v", got, want)
	}
}
