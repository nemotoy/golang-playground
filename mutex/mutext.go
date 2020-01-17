package main

import (
	"fmt"
	"strconv"
	"sync"
)

type p struct {
	s []string
	i []int
	sync.RWMutex
}

const max = 100

func main() {
	p := p{}
	go func() {
		for i := 0; i < max; i++ {
			p.Lock()
			fmt.Printf("#S: %d\n", i)
			p.s = append(p.s, strconv.Itoa(i))
			p.Unlock()
		}
	}()
	for i := 0; i < max; i++ {
		p.Lock()
		fmt.Printf("#I: %d\n", i)
		p.i = append(p.i, i)
		p.Unlock()
	}
	p.Lock()
	fmt.Printf("#I: %v\n", p.i)
	p.Unlock()
	p.Lock()
	fmt.Printf("#S: %v\n", p.s)
	p.Unlock()
}
