package flag

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

type opt struct {
	s string
	i int
	b bool
}

func (o *opt) parse(args []string) error {
	f := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	f.StringVar(&o.s, "n", "", "string flag")
	f.IntVar(&o.i, "a", 0, "int flag")
	f.BoolVar(&o.b, "b", false, "bool flag")
	return f.Parse(args)
}

func Test_parse(t *testing.T) {
	tests := []struct {
		name    string
		in      []string
		want    *opt
		wantErr bool
	}{
		{
			name:    "correct args",
			in:      []string{"-n", "sss", "-a", "100", "-b", "true"},
			want:    &opt{s: "sss", i: 100, b: true},
			wantErr: false,
		},
		{
			name:    "empty args",
			in:      []string{},
			want:    &opt{s: "", i: 0, b: false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var o = &opt{}
			err := o.parse(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(o, tt.want) {
				t.Errorf("parse() got = %+v, but want = %+v", o, tt.want)
			}
		})
	}
}
