package roundtrip

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
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

func Test_GetReqBody(t *testing.T) {
	v := user{"test"}
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	in, err := http.NewRequest("POST", "test", bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	got, err := GetReqBody(in)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, b) {
		t.Errorf("parse() got = %+v, but want = %+v", got, b)
	}
}

func GetReqBody(req *http.Request) ([]byte, error) {
	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, errors.New("request body is nil")
}

func CopyReq(req *http.Request) (*http.Request, error) {
	if req.GetBody != nil {
		newReq := *req
		var err error
		newReq.Body, err = req.GetBody()
		if err != nil {
			return nil, err
		}
		req = &newReq
		return req, nil
	}
	return nil, errors.New("request body is nil")
}

func Benchmark_GetReqBody(b *testing.B) {
	v := user{"test"}
	body, err := json.Marshal(v)
	if err != nil {
		b.Fatal(err)
	}
	in, err := http.NewRequest("POST", "test", bytes.NewReader(body))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if outBody, err := GetReqBody(in); err != nil {
			b.Log(outBody, err)
		}
	}
}

func Benchmark_CopyReq(b *testing.B) {
	v := user{"test"}
	body, err := json.Marshal(v)
	if err != nil {
		b.Fatal(err)
	}
	in, err := http.NewRequest("POST", "test", bytes.NewReader(body))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if req, err := CopyReq(in); err != nil {
			b.Log(req, err)
		}
	}
}
