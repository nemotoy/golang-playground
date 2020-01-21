package retry

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"context"

	"github.com/cenkalti/backoff"
)

var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
})

func Test_(t *testing.T) {
	ts := httptest.NewServer(handler)
	defer ts.Close()

	c := ts.Client()
	var r *http.Response
	ctx := context.Background()
	err := backoff.Retry(func() (err error) {
		r, err = c.Get(ts.URL)
		if err == nil && r.StatusCode == http.StatusOK {
			return nil
		}
		r.Body.Close()
		return err
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
	if err != nil {
		t.Fatal(err)
	}
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
