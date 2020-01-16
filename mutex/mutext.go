package main

import (
	"fmt"
	"sync"
)

type p struct {
	i []int
	sync.RWMutex
}

const max = 100

func main() {
	p := p{}
	for i := 0; i < max; i++ {
		p.Lock()
		p.i = append(p.i, i)
		p.Unlock()
	}
	fmt.Println(p.i)
}
