package greet

import "fmt"

type user struct {
	name string
}

func greet() string {
	return "hello, world"
}

func (u *user) greet() string {
	return fmt.Sprintf("I am %s", u.name)
}
