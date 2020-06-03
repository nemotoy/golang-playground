package reflect

import (
	"reflect"
)

const (
	expected = "test"
)

func isStringValWithReflect(v interface{}) bool {
	switch v := reflect.ValueOf(v); v.Kind() {
	case reflect.String:
		if v.String() == expected {
			return true
		}
	}
	return false
}

func isStringValWithTypeAssert(i interface{}) bool {
	if v, ok := i.(string); ok {
		return v == expected
	}
	return false
}
