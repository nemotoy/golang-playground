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
	return fmt.Sprintf("{id: %d, name: %s, body: %s}", o.id, o.name, string(o.body))
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
		// Note: spew prints the struct name (e.g. (print.<obuject name>))
		if strings.Contains(want, got) {
			t.Errorf("got = %s, but want = %s", got, want)
		}
	})
}
