package examples

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.Methods("GET").Path("/ping").HandlerFunc(AuthMiddleware(ping))
	r.Methods("GET").Path("/user").Handler(new(userImpl))
	return r
}

// TODO: replace signature to func(f http.Handler) http.Handler
var AuthMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if v := r.Header.Get("X-Auth-Id"); v == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		f(w, r)
	}
}

func TestHandler(t *testing.T) {
	handler := initHandler()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	t.Run("ping", func(t *testing.T) {
		{
			e.Builder(func(req *httpexpect.Request) {
				req.WithHeader("X-Auth-Id", "test")
			}).GET("/ping").
				Expect().
				Status(http.StatusOK).ContentType("text/plain").Text().Equal("pong\n")
		}
		{
			e.GET("/ping").
				Expect().
				Status(http.StatusUnauthorized)
		}
	})
	t.Run("user", func(t *testing.T) {
		{
			raw := e.GET("/user").
				Expect().
				Status(http.StatusOK).ContentType("application/json").JSON().Object()
			raw.ContainsMap(map[string]interface{}{
				"name": "hoge",
			})
		}
	})
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
