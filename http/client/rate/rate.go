package main

import (
	"context"
	"log"
	"os"
	"sync"
)

// rf. ratelimit in https://github.com/kat-co/concurrency-in-go-src
func main() {
	defer log.Printf("Done.")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConnection := Open()
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if err := apiConnection.ReadFile(context.Background()); err != nil {
				log.Printf("cannot ReadFile: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if err := apiConnection.ResolveAddress(context.Background()); err != nil {
				log.Printf("cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	wg.Wait()
}

type APIConnection struct{}

func Open() *APIConnection {
	return &APIConnection{}
}

func (a *APIConnection) ReadFile(ctx context.Context) error { return nil }

func (a *APIConnection) ResolveAddress(ctx context.Context) error { return nil }
