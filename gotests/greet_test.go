package greet

import (
	"sync"
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

func Test_user_getM(t *testing.T) {
	type fields struct {
		name string
		m    *sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      string
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				name: tt.fields.name,
				m:    tt.fields.m,
			}
			got, err := u.getM(tt.args.key)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
