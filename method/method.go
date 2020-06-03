package method

import "errors"

type User struct {
	Name string
}

var (
	errReceiverNil = errors.New("receiver is nil")
	errNameEmpty   = errors.New("given name is empty")
)

func (u *User) updateName(name string) error {
	if u == nil {
		return errReceiverNil
	}
	if name == "" {
		return errNameEmpty
	}
	u.Name = name
	return nil
}
