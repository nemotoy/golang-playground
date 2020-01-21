package retry

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/*
https://github.com/cenkalti/backoff

*/

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
})

func Test_(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	c := ts.Client()
	r, err := c.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	got := string(b)
	want := "hello"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %v; want %v", got, want)
	}
}
