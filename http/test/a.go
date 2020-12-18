package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("pong"))
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
	authMiddleware := &authMiddleware{}
	r.Methods("GET").Path("/ping").HandlerFunc(ping)
	r.Methods("GET").Path("/user").Handler(userHandler)
	r.Methods("GET").Path("/user/{id:[0-9]+}").Handler(userHandler)
	r.Methods("GET").Path("/user/{name}").Handler(userHandler)
	r.Use(authMiddleware.Middleware)
	return r
}

type authMiddleware struct{}

func (am *authMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := r.Header.Get("X-Auth-Id")
		if v == "" {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var LoggerMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: implment
		f(w, r)
	}
}

func main() {
	handler := initHandler()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Println("failed to serve: ", err)
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, os.Interrupt)
	<-s
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("failed to shutdown gracefully: ", err)
	}
	log.Println("Server shutdown")

}
