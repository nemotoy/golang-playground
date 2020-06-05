package roundtrip

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"golang.org/x/net/context"
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

type user struct {
	Name string `json:"name"`
}

func Test_RoundtripWithPOST(t *testing.T) {
	counter := 0
	wantRetries := 3
	wantName := "hoge"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Received a request: %+v", r)
		v := &user{}
		if err := json.NewDecoder(r.Body).Decode(v); err != nil || v.Name != wantName {
			t.Fatal(err)
		}
		t.Logf("Read body form a request: %v", v)
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
	v := user{wantName}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", ts.URL, bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	ctx, cf := context.WithTimeout(context.Background(), 10*time.Second)
	defer cf()
	_ = req.WithContext(ctx)
	resp, err := c.Do(req)
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
