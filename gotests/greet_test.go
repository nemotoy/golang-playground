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
