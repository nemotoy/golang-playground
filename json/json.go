package json

import "encoding/json"

type Sample struct {
	Key  string `json:"key"`
	Data Data   `json:"data"`
}

type Data interface {
	Data()
}

type SourceA struct {
	S string `json:"s"`
	I int    `json:"i,string"`
	B bool   `json:"b,string"`
}

func (*SourceA) Data() {}

func (s *Sample) UnmarshalJSON(data []byte) error {
	type alias Sample
	a := struct {
		Data json.RawMessage `json:"data"`
		*alias
	}{
		alias: (*alias)(s),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	// store an implementation of the interface
	var sa SourceA
	if err := json.Unmarshal(a.Data, &sa); err != nil {
		return err
	}
	// &sa implement the Data interface, so could assign it.
	s.Data = &sa

	return nil
}
