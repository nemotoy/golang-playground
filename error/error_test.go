package error

import (
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
