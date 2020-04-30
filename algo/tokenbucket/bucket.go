package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var defaultInterval = 10 * time.Second

const tokenValue = "t"

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
	go b.pourTokens(1*time.Minute, 10)

	// sample to access the bucket
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()
	for {
		<-t.C
		log.Println("elapsed a given interval")
		b.mu.Lock()
		// If succeed to get it, delete it and pass the next step
		log.Println("access the bucket to get a token")
		if len(b.tokens) == 0 {
			log.Println("so a token doesn't exist in the bucket, sleep a few miliseconds")
			time.Sleep(500 * time.Millisecond)
			b.mu.Unlock()
			continue
		}
		b.tokens = append(b.tokens[:0], b.tokens[1:]...)
		log.Println("tokens: ", len(b.tokens))
		b.mu.Unlock()
		log.Println("do something")
	}
}

func (b *bucket) setup() error {
	if b == nil {
		return fmt.Errorf("bucket is empty")
	}
	for i := 1; i <= b.size; i++ {
		b.tokens = append(b.tokens, tokenValue)
	}
	return nil
}

// pourToken adds tokens to the bucket at a given interval
func (b *bucket) pourTokens(interval time.Duration, tokenNum int) {
	if interval < 0 {
		interval = defaultInterval
	}
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		<-t.C
		b.mu.Lock()
		log.Println("pour tokens")
		for i := 1; i <= tokenNum; i++ {
			if len(b.tokens) == b.size {
				log.Println("so the bucket's capacity is filled, stop pouring tokens")
				break
			}
			b.tokens = append(b.tokens, tokenValue)
		}
		log.Println("current number of tokens: ", len(b.tokens))
		b.mu.Unlock()
	}
}
