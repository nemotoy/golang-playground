package hash

import (
	"fmt"

	"github.com/mitchellh/hashstructure"
)

func genHash(v interface{}) (string, error) {
	hash, err := hashstructure.Hash(v, nil)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash), nil
}
