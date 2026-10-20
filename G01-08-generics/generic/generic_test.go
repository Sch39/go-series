package generic_test

import (
	"strconv"
	"testing"

	"github.com/sch39/go-series/g0108/generics/generic"
)

type Score int

func TestContainsAndSum(t *testing.T) {
	values := []Score{10, 20, 30}

	found := generic.Contains(values, Score(20))
	total := generic.Sum(values)
	t.Logf("found=%v total=%d", found, total)

	if !found {
		t.Error("Contains returned false, want true")
	}
	if total != 60 {
		t.Errorf("Sum returned %d, want 60", total)
	}
}

func TestMapChangesElementType(t *testing.T) {
	values := []int{10, 20, 30}
	strings := generic.Map(values, strconv.Itoa)
	t.Logf("values=%v strings=%v", values, strings)

	if len(strings) != 3 || strings[0] != "10" || strings[2] != "30" {
		t.Errorf("Map returned %v, want [10 20 30]", strings)
	}
}

func TestStack(t *testing.T) {
	var stack generic.Stack[string]
	stack.Push("first")
	stack.Push("second")

	value, ok := stack.Pop()
	t.Logf("value=%q ok=%v", value, ok)
	if !ok || value != "second" {
		t.Errorf("Pop returned %q, %v; want second, true", value, ok)
	}

	stack.Pop()
	zero, ok := stack.Pop()
	t.Logf("empty value=%q ok=%v", zero, ok)
	if ok || zero != "" {
		t.Errorf("empty Pop returned %q, %v; want empty string, false", zero, ok)
	}
}
