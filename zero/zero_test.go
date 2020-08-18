package main

import (
	"reflect"
	"testing"
)

type stub struct {
	s string
}

func Test_zero(t *testing.T) {
	var ii map[int64]stub
	t.Logf("value: %#v, len: %d", ii, len(ii))
	for _, i := range ii {
		t.Log(i)
	}

	ii2 := make(map[int64]stub, 3)
	t.Logf("value: %#v, len: %d", ii2, len(ii2))
	for _, i2 := range ii2 {
		t.Log(i2)
	}

	var ii3 = []int64{}
	t.Log(len(ii3))

	for _, v := range ii3 {
		t.Log(v)
	}

	var ss []string
	t.Logf("value: %#v, len: %d", ss, len(ss))
	for i, s := range ss {
		t.Log(i, s)
	}
	ss2 := make([]string, 5)
	t.Logf("value: %#v, len: %d", ss2, len(ss2))
	for i, s2 := range ss2 {
		t.Log(i, s2)
	}
	if reflect.DeepEqual(ss, ss2) {
		t.Errorf("got = %+v, but want = %+v", ss, ss2)
	}
}
