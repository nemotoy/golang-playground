package greet

import (
	"fmt"
	"sync"
)

type user struct {
	name string
	m    *sync.Map
}

func greet() string {
	return "hello, world"
}

func (u *user) greet() string {
	return fmt.Sprintf("I am %s", u.name)
}

func (u *user) getM(key string) (string, error) {
	raw, ok := u.m.Load(key)
	if !ok {
		return "", fmt.Errorf("")
	}
	v, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("")
	}
	return v, nil
}
