package reflect

import (
	"reflect"
)

const (
	expected = "test"
)

// nolint:unused
func isStringVal(v interface{}) bool {
	switch v := reflect.ValueOf(v); v.Kind() {
	case reflect.String:
		if v.String() == expected {
			return true
		}
	}
	return false
}
