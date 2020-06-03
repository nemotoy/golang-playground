package reflect

import "testing"

func Test_is(t *testing.T) {
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
			if got := isStringVal(tt.args.v); got != tt.want {
				t.Errorf("isStringVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
