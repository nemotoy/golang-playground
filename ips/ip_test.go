package ips

import (
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
		{"192.0000.0000.0000", false},
		{"192.168.5.1", true},
		{"127.0.0.1/32", true},
		{"0000:0000:0000:0000:0000:ffff:7f00:0001", true},
	}
	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			if isIP(tt.ip) == tt.want {
				t.Error("unexpeccted")
			}
		})
	}
}

func isIP(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		_, _, err := net.ParseCIDR(s)
		return err != nil
	}
	return true
}
