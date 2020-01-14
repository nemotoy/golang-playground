package main

import (
	"fmt"
	"time"
)

type m map[int]string

func main() {
	m := make(m)
	m[1] = "1"
	m[2] = "2"
	go func() {
		delete(m, 1)
	}()
	v := m[1]
	fmt.Println(v)
	time.Sleep(3 * time.Second)
}
