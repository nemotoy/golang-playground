package print

import (
	"fmt"
	"testing"
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
	got := o.String()
	if got != want {
		t.Fail()
	}
}
