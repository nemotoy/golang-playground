package flag

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

type opt struct {
	name string
}

func parse(args []string) (opt, error) {
	var o = opt{}
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&o.name, "n", "", "name")
	return o, f.Parse(args)
}

func Test_parse(t *testing.T) {
	in := []string{"-n", "nameA"}
	want := opt{name: "nameA"}
	got, err := parse(in)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got = %+v, but want = %+v", got, want)

	}
}
