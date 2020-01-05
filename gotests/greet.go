package greet

import (
	"fmt"
)

type user struct {
	name string
}

func greet() string {
	return "hello, world"
}

func (u *user) greet() string {
	return fmt.Sprintf("I am %s", u.name)
}

const (
	max = 100
	min = 1
)

func cal(n int) (string, error) {
	if n > max {
		return "", fmt.Errorf("over the max limit(%d); given: %d", max, n)
	}
	return fmt.Sprintf("%d", n), nil
}

func cal2(n int) (string, error) {
	if n > max {
		return "", fmt.Errorf("over the max limit(%d); given: %d", max, n)
	}
	return fmt.Sprintf("%d", n), nil
}
