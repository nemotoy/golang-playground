package ips

import (
	"net"
	"testing"
)

func TestValidateIp(t *testing.T) {
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
			if isIP2(tt.ip) != tt.want {
				t.Error("unexpeccted")
			}
		})
	}
}

func isIP(s string) bool {
	_, _, err := net.ParseCIDR(s)
	if err != nil {
		ip := net.ParseIP(s)
		if ip == nil {
			return false
		}
	}
	return true
}

func isIP2(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		_, _, err := net.ParseCIDR(s)
		if err != nil {
			return false
		}
	}
	return true
}

func Benchmark_IsIP(b *testing.B) {
	s := "127.0.0.1/32"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isIP(s)
	}
}

func Benchmark_IsIP2(b *testing.B) {
	s := "127.0.0.1/32"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = isIP2(s)
	}
}
