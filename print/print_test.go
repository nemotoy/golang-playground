package print

import (
	"fmt"
	"testing"

	"strings"

	"github.com/davecgh/go-spew/spew"
)

type o struct {
	id   int
	name string
	body []byte
}

func (o o) String() string {
	return fmt.Sprintf("{id: %d, name: %s, body: %s}", o.id, o.name, o.body)
}

func Test_print(t *testing.T) {
	want := "{id: 100, name: A, body: {t: ttt}}"
	o := o{
		id:   100,
		name: "A",
		body: []byte("{t: ttt}"),
	}
	t.Run("Use String method", func(t *testing.T) {
		got := o.String()
		if got != want {
			t.Errorf("got = %s, but want = %s", got, want)
		}
	})
	t.Run("Use davecgh/go-spew/spew", func(t *testing.T) {
		got := spew.Sdump(o)
		// Note: spew prints the struct name (e.g. (print.<object name>))
		if strings.Contains(want, got) {
			t.Errorf("got = %s, but want = %s", got, want)
		}
	})
}

/*
	$ go test -count 3 -benchmem -bench .
	goos: darwin
	goarch: amd64
	pkg: github.com/nemotoy/golang-playground/print
	Benchmark_String-4       2664465               448 ns/op              96 B/op          5 allocs/op
	Benchmark_String-4       2392327               558 ns/op              96 B/op          5 allocs/op
	Benchmark_String-4       2574836               465 ns/op              96 B/op          5 allocs/op
	Benchmark_Sdump-4         721947              1397 ns/op             408 B/op         12 allocs/op
	Benchmark_Sdump-4         846585              1372 ns/op             408 B/op         12 allocs/op
	Benchmark_Sdump-4         919640              1377 ns/op             408 B/op         12 allocs/op
	PASS
	ok      github.com/nemotoy/golang-playground/print      11.500s
*/
func Benchmark_String(b *testing.B) {
	o := o{
		id:   100,
		name: "A",
		body: []byte("{t: ttt}"),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = o.String()
	}
}

func Benchmark_Sdump(b *testing.B) {
	o := o{
		id:   100,
		name: "A",
		body: []byte("{t: ttt}"),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = spew.Sdump(o)
	}
}
