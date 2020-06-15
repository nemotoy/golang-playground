package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(rate.Limit(1), 1)
	router := mux.NewRouter()
	router.Methods(http.MethodGet).Path("/ping").Handler(&PingImpl{limiter})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
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

type PingImpl struct {
	limiter *rate.Limiter
}

func (p *PingImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Received")
	ctx := context.Background()
	if err := p.limiter.Wait(ctx); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("ping"))
}
