package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	m := &sync.Map{}
	m.Store(1, "1")
	m.Store(2, "2")
	go func() {
		m.Delete(1)
	}()
	v, ok := m.Load(1)
	if !ok {
		fmt.Fprintf(os.Stderr, "failed to load")
		return
	}
	fmt.Println(v)
	time.Sleep(3 * time.Second)
}
