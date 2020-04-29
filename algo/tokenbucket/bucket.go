package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type bucket struct {
	mu     sync.Mutex
	tokens []string
	size   int
}

// learn the token bucket
// https://ja.wikipedia.org/wiki/%E3%83%88%E3%83%BC%E3%82%AF%E3%83%B3%E3%83%90%E3%82%B1%E3%83%83%E3%83%88
func main() {
	// setup a bucket
	b := &bucket{
		size: 100,
	}
	if err := b.setup(); err != nil {
		log.Fatal(err)
	}
	// start adding tokens
	go b.pourTokens(1 * time.Second)

	// sample to access the bucket
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()
	for {
		<-t.C
		log.Println("elapsed a given interval")
		b.mu.Lock()
		// If succeed to get it, delete it and pass the next step
		log.Println("access the bucket to get a token")
		b.mu.Unlock()
	}
}

func (b *bucket) setup() error {
	if b == nil {
		return fmt.Errorf("bucket is empty")
	}
	for i := 1; i <= b.size; i++ {
		log.Printf("i: %d\n", i)
		b.tokens = append(b.tokens, "t")
	}
	log.Printf("amount of tokens are %d\n", len(b.tokens))
	return nil
}

// todo: gives pouring number of token per interval
// pourToken adds tokens to the bucket at a given interval
func (b *bucket) pourTokens(interval time.Duration) {
	if interval < 0 {
		interval = 10 * time.Second
	}
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		<-t.C
		b.mu.Lock()
		log.Println("pour tokens")
		b.mu.Unlock()
	}
}
