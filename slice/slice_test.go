package slice

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_f(t *testing.T) {
	in := []string{"a", "b", "c"}
	f(in)
	want := []string{"a", "b", "c", "d"}
	if diff := cmp.Diff(in, want); diff == "" {
		t.Errorf("Find() mismatch (-want +got):\n%s", diff)
	}
}

func Test_f2(t *testing.T) {
	in := []string{"a", "b", "c"}
	f2(&in)
	want := []string{"a", "b", "c", "d"}
	if diff := cmp.Diff(in, want); diff != "" {
		t.Errorf("Find() mismatch (-want +got):\n%s", diff)
	}
}

var f = func(ss []string) {
	ss = append(ss, "d")
}

var f2 = func(ss *[]string) {
	*ss = append(*ss, "d")
}
