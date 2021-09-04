package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nemotoy/golang-playground/ddd/server/application"
	"github.com/nemotoy/golang-playground/ddd/server/handler"
)

func main() {
	mux := http.NewServeMux()
	userAppSrv := application.NewUserApplicationService()
	userHandler := handler.NewUserHandler(userAppSrv)

	mux.HandleFunc("/ping", ping)
	mux.Handle("/users", userHandler)
	srv := &http.Server{
		Addr:    ":8085",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Println("failed to serve: ", err)
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM, os.Interrupt)
	<-s
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("failed to shutdown gracefully: ", err)
	}
	log.Println("Server shutdown")
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received Request: %v\n", r)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("pong\n"))
}
