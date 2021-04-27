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
		{"192.168.5.1", true},
		{"127.0.0.1/32", true},
		{"0000:0000:0000:0000:0000:ffff:7f00:0001", true},
	}
	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			ip := net.ParseIP(tt.ip)
			if ip == nil && tt.want {
				ip, ipNet, err := net.ParseCIDR(tt.ip)
				if err != nil && !tt.want {
					t.Error("unexpeccted")
				}
				t.Log(ip, ipNet)
			}
		})
	}
}
