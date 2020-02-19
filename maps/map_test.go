package maps

import (
	"strconv"
	"testing"
)

type o struct {
	id  int
	key string
}

/*
	$ go test -count 3 -benchmem -bench .
	goos: darwin
	goarch: amd64
	pkg: github.com/nemotoy/golang-playground/maps
	Benchmark_nocaps-4       4188104               401 ns/op             155 B/op          0 allocs/op
	Benchmark_nocaps-4       3853094               443 ns/op             168 B/op          0 allocs/op
	Benchmark_nocaps-4       3459174               343 ns/op             188 B/op          0 allocs/op
	Benchmark_caps-4        11782938               156 ns/op               2 B/op          0 allocs/op
	Benchmark_caps-4        10997864               146 ns/op               1 B/op          0 allocs/op
	Benchmark_caps-4        10179909               137 ns/op               0 B/op          0 allocs/op
	PASS
	ok      github.com/nemotoy/golang-playground/maps       23.879s
*/
func Benchmark_nocaps(b *testing.B) {
	n := b.N
	var oo = []o{}
	for i := 0; i < n; i++ {
		o := o{id: i, key: strconv.Itoa(i)}
		oo = append(oo, o)
	}
	om := make(map[int]o)
	b.ResetTimer()
	for i := 0; i < n; i++ {
		om[i] = oo[i]
	}
}

func Benchmark_caps(b *testing.B) {
	n := b.N
	var oo = []o{}
	for i := 0; i < n; i++ {
		o := o{id: i, key: strconv.Itoa(i)}
		oo = append(oo, o)
	}
	om := make(map[int]o, n)
	b.ResetTimer()
	for i := 0; i < n; i++ {
		om[i] = oo[i]
	}
}
