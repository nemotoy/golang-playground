package json

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_SampleUnmarshalJSON(t *testing.T) {
	in := `{"key":"keyA", "data":[{"b":"false","i":"0","s":""}]}`
	out := &Sample{
		Key: "keyA",
		Data: []Data{
			&SourceA{
				S: "SSS",
				I: 100,
				B: true,
			},
		},
	}
	var v Sample
	if err := json.Unmarshal([]byte(in), &v); err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(v, out); diff != "" {
		t.Errorf("Find() mismatch (-want +got):\n%s", diff)
	}
}

type Sample struct {
	Key  string `json:"key"`
	Data []Data `json:"data"`
}

type Data interface {
	Data()
}

type SourceA struct {
	S string `json:"s"`
	I int    `json:"i"`
	B bool   `json:"b"`
}

func (*SourceA) Data() {}
