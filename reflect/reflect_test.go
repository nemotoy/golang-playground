package reflect

import (
	"testing"
)

func Test_isStringValWithReflect(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"type of string, expected value", args{v: expected}, true},
		{"type of string, unexpected value", args{v: "aaaa"}, false},
		{"type of int", args{v: 100}, false},
		{"type of struct", args{v: struct{}{}}, false},
		{"type of string slice", args{v: []string{expected}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStringValWithReflect(tt.args.v); got != tt.want {
				t.Errorf("isStringValWithReflect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStringValWithTypeAssert(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"type of string, expected value", args{i: expected}, true},
		{"type of string, unexpected value", args{i: "aaaa"}, false},
		{"type of int", args{i: 100}, false},
		{"type of struct", args{i: struct{}{}}, false},
		{"type of string slice", args{i: []string{expected}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStringValWithTypeAssert(tt.args.i); got != tt.want {
				t.Errorf("isStringValWithTypeAssert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isStringValWithReflect(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isStringValWithReflect(expected)
	}
}

func Benchmark_isStringValWithTypeAssert(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isStringValWithTypeAssert(expected)
	}
}
