package rebase

import (
	"testing"
)

func TestIsBaseValid(t *testing.T) {
	validPairs := []struct {
		base  int
		valid bool
	}{
		{1, false},
		{23, true},
		{63, false},
	}
	for _, p := range validPairs {
		got := IsBaseValid(p.base)
		if got != p.valid {
			t.Errorf("IsBaseValid(%d) = %q, want %q", p.base, got, p.valid)
		}
	}
}
