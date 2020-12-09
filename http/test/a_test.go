package examples

import (
	"encoding/json"
	"fmt"
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

var stubUsers = map[string]User{
	"aaa": {Name: "aaa"},
	"bbb": {Name: "bbb"},
	"ccc": {Name: "ccc"},
}

type userImpl struct {
	users map[string]User // TODO: replace it with collecting from a DB system.
}

type User struct {
	Name string `json:"name"`
}

func (u *userImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	user, ok := u.users[vars["name"]]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(user)
	_, _ = w.Write(b)
}

type Logger interface {
	Printf(format string, v ...interface{})
}

func initHandler() http.Handler {
	r := mux.NewRouter()
	userHandler := &userImpl{stubUsers}
	r.Methods("GET").Path("/ping").HandlerFunc(LoggerMiddleware(AuthMiddleware(ping)))
	r.Methods("GET").Path("/user").Handler(userHandler)
	r.Methods("GET").Path("/user/{id:[0-9]+}").Handler(userHandler)
	r.Methods("GET").Path("/user/{name}").Handler(userHandler)
	return r
}

// TODO: replace signature to func(f http.Handler) http.Handler
var AuthMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := r.Header.Get("X-Auth-Id")
		if v == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Printf("Auth-id: %s\n", v)
		f(w, r)
	}
}

var LoggerMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: implment
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
		{
			e.POST("/ping").
				Expect().
				Status(http.StatusMethodNotAllowed)
		}
	})
	t.Run("user", func(t *testing.T) {
		{
			e.GET("/user").
				Expect().
				Status(http.StatusNotFound)
		}
		{
			e.GET("/user/111").
				Expect().
				Status(http.StatusNotFound)
		}
		{
			raw := e.GET("/user/aaa").
				Expect().
				Status(http.StatusOK).ContentType("application/json").JSON().Object()
			raw.ContainsMap(map[string]interface{}{
				"name": "aaa",
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
