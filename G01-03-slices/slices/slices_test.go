package slices_test

import (
	"testing"

	"github.com/sch39/go-series/g0103/slices"
)

func TestLengthCapacityAndAppend(t *testing.T) {
	values := make([]int, 2, 4)
	length, capacity := slices.Describe(values)
	if length != 2 || capacity != 4 {
		t.Fatalf("Describe() = len %d cap %d, want len 2 cap 4", length, capacity)
	}

	grown := slices.AppendOne(values, 7)
	t.Logf("before len=%d cap=%d; after=%v len=%d cap=%d", length, capacity, grown, len(grown), cap(grown))
	if len(grown) != 3 || grown[2] != 7 {
		t.Errorf("AppendOne() = %v, want third element 7", grown)
	}
}

func TestAliasingAndIsolatedAppend(t *testing.T) {
	values := make([]int, 2, 3)
	values[0] = 10
	alias := values[:1]
	slices.SetFirst(alias, 99)
	t.Logf("values after alias mutation=%v", values)
	if values[0] != 99 {
		t.Errorf("SetFirst() did not mutate shared backing array; got %v", values)
	}

	isolated := slices.IsolateAppend(values[:1], 42)
	t.Logf("original=%v isolated=%v", values, isolated)
	if len(isolated) != 2 || isolated[1] != 42 {
		t.Errorf("IsolateAppend() = %v, want [99 42]", isolated)
	}
	if len(values) != 2 || values[1] != 0 {
		t.Errorf("IsolateAppend() changed original length/data; got %v", values)
	}
}
