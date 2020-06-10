package error

import (
	"errors"
	"fmt"
	"testing"

	"go.uber.org/multierr"
)

// practice https://github.com/uber-go/multierr
func Test_mltierr(t *testing.T) {
	var (
		e1 = fmt.Errorf("e1")
		e2 = fmt.Errorf("e2")
		e3 = fmt.Errorf("e3")
	)

	var ae error
	ae = multierr.Append(ae, e1)
	ae = multierr.Append(ae, e2)
	ae = multierr.Append(ae, e3)
	t.Log(ae)

	var aie error
	isAppended := multierr.AppendInto(&aie, e1)
	t.Log(isAppended)

	ce := multierr.Combine(e1, e2, e3)
	t.Log(ce)

	ee := multierr.Errors(ce)
	t.Log(ee)
}

func Benchmark_IsExErr(b *testing.B) {
	err := &ExError{"one", errors.New("one")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if ok := IsExErr(err); ok {
			b.Log(ok)
		}
	}
}

func Benchmark_ClassifyErr(b *testing.B) {
	err := &ExError{"one", errors.New("one")}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := ClassifyErr(err); err != nil {
			b.Log(err)
		}
	}
}
