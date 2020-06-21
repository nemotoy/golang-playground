package json

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_SampleUnmarshalJSON(t *testing.T) {
	in := `{"key":"keyA","data":{"s":"SSS","i":"100","b":"true"}}`
	out := Sample{
		Key: "keyA",
		Data: &SourceA{
			S: "SSS",
			I: 100,
			B: true,
		},
	}
	var v Sample
	if err := json.Unmarshal([]byte(in), &v); err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(out, v); diff != "" {
		t.Errorf("Find() mismatch (-want +got):\n%s", diff)
	}
}

func Test_SampleMarshalJSON(t *testing.T) {
	in := Sample{
		Key: "keyA",
		Data: &SourceA{
			S: "SSS",
			I: 100,
			B: true,
		},
	}
	out := `{"key":"keyA","data":{"s":"SSS","i":"100","b":"true"}}`
	v, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff([]byte(out), v); diff != "" {
		t.Errorf("Find() mismatch (-want +got):\n%s", diff)
	}
}

func Benchmark_SampleUnmarshalJSON(b *testing.B) {
	in := `{"key":"keyA","data":{"s":"SSS","i":"100","b":"true"}}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v Sample
		if err := json.Unmarshal([]byte(in), &v); err != nil {
			b.Fatal(err)
		}
	}
}
