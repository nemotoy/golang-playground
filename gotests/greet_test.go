package greet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_greet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greet(); got != tt.want {
				t.Errorf("greet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_user_greet(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				name: tt.fields.name,
			}
			assert.Equal(t, tt.want, u.greet())
		})
	}
}

func Test_cal(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		want      string
		assertion assert.ErrorAssertionFunc
	}{
		{
			name:      "",
			args:      args{n: 0},
			want:      "0",
			assertion: assert.NoError,
		},
		{
			name:      "",
			args:      args{n: 101},
			want:      "",
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cal(tt.args.n)
			tt.assertion(t, err)
			// 自動生成されるのはEqual関数．
			assert.Equal(t, tt.want, got)
		})
	}
}
