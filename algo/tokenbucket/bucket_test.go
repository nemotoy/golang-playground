package main

import "testing"

func Test_setup(t *testing.T) {
	tests := []struct {
		inSize int
	}{
		{0},
		{100},
		{1000},
		{10000},
	}
	for _, test := range tests {
		b := &bucket{
			size: test.inSize,
		}
		if err := b.setup(); err != nil {
			t.Fatal(err)
		}
		if len(b.tokens) != test.inSize {
			t.Errorf("unexpected number of tokens; got: %d, want: %d", len(b.tokens), test.inSize)
		}
	}
}
