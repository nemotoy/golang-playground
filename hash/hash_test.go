package hash

import "testing"

func Test_genHash(t *testing.T) {

	t.Run("unordered slices", func(t *testing.T) {
		v1 := []string{"a", "b", "c"}
		v2 := []string{"c", "a", "b"}
		wantErr := false

		got, err := genHash(v1)
		if (err != nil) != wantErr {
			t.Errorf("genHash() error = %v, wantErr %v", err, wantErr)
			return
		}
		got2, err := genHash(v2)
		if (err != nil) != wantErr {
			t.Errorf("genHash() error = %v, wantErr %v", err, wantErr)
			return
		}
		if got == got2 {
			t.Errorf("genHash() = %v, want %v", got, got2)
		}
	})
}
