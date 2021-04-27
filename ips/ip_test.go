package ips

import (
	"fmt"
	"net"
	"testing"
)

func TestParseIp(t *testing.T) {
	tests := []struct {
		ip   string
		want bool
	}{
		{"test", false},
		{"test/test", false},
		{"test.test.test", false},
		{"192.168.5.1", true},
		{"127.0.0.1/32", true},
		{"2001:db8::/32", true},
	}
	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			if isIP(tt.ip) != tt.want {
				t.Error("unexpeccted")
			}
		})
	}
}

func isIP(s string) bool {
	_, _, err := net.ParseCIDR(s)
	if err != nil {
		fmt.Println(s)
		ip := net.ParseIP(s)
		if ip == nil {
			return false
		}
	}
	return true
}
