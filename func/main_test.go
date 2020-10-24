package main

import (
	"fmt"
	"testing"
)

type f func() string

func Test_xxx(t *testing.T) {

	f2 := func() string {
		return "hello"
	}
	fmt.Println("hello", f2())

	hoge := &hoge{}
	f2 = hoge.Get

	fmt.Printf("expected = hoge, got = %s\n", f2())
}

type hoge struct{}

func (h *hoge) Get() string {
	return "hoge"
}
