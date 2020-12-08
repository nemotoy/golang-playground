package examples

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("pong\n"))
}

type userImpl struct {
}

type User struct {
	Name string `json:"name"`
}

func (u *userImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := User{Name: "hoge"}
	b, _ := json.Marshal(res)
	_, _ = w.Write(b)
}

func initHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		ping(w, r)
	})
	mux.Handle("/user", new(userImpl))
	return mux
}

func TestPing(t *testing.T) {
	handler := initHandler()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	{
		e.GET("/ping").
			Expect().
			Status(http.StatusOK).ContentType("text/plain").Text().Equal("pong\n")
	}
	{
		raw := e.GET("/user").
			Expect().
			Status(http.StatusOK).ContentType("application/json").JSON().Object()
		raw.ContainsMap(map[string]interface{}{
			"name": "hoge",
		})
	}
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() {
}

func teardown() {
}
