package model

import (
	"fmt"
)

type User struct {
	LastName  LastName
	FirstName FirstName
}

type LastName string

func (ln LastName) satisfiedLength() bool {
	l := len(ln)
	return 0 < l && l <= 128
}

type FirstName string

func (ln FirstName) satisfiedLength() bool {
	l := len(ln)
	return 0 < l && l <= 128
}

func NewUser(lastName LastName, firstName FirstName) (*User, error) {
	if !lastName.satisfiedLength() {
		return nil, fmt.Errorf("")
	}
	if !firstName.satisfiedLength() {
		return nil, fmt.Errorf("")
	}
	return &User{
		LastName:  lastName,
		FirstName: firstName,
	}, nil
}
