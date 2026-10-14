package data_test

import (
	"testing"

	"github.com/sch39/go-series/g0102/valueptr/data"
)

func TestShallowCopy(t *testing.T) {
	a := data.Data{
		Tags:   map[string]bool{"go": true},
		Scores: []int{10, 20},
	}
	b := a

	// Map tetap dibagi
	b.Tags["x"] = true
	if !a.Tags["x"] {
		t.Errorf("map mutation should be visible in a; a.Tags=%v", a.Tags)
	}

	// Slice elemen tetap dibagi
	b.ChangeFirstScore()
	if a.Scores[0] != 99 {
		t.Errorf("slice element mutation should be visible in a; got %d, want 99", a.Scores[0])
	}

	// Slice append tidak terlihat
	lenBefore := len(a.Scores)
	b.AppendScore()
	if len(a.Scores) != lenBefore {
		t.Errorf("slice append in copy should not be visible in a; got len=%d, want %d", len(a.Scores), lenBefore)
	}
}
