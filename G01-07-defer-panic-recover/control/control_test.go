package control_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/sch39/go-series/g0107/control/control"
)

func TestDeferOrder(t *testing.T) {
	got := control.CleanupOrder()
	want := []string{"work", "second", "first"}

	t.Logf("events=%v", got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("CleanupOrder() = %v, want %v", got, want)
	}
}

func TestDeferredArgumentsAreEvaluatedImmediately(t *testing.T) {
	got := control.EvaluatedArgument()

	t.Logf("captured=%q", got)

	if got != "before" {
		t.Errorf("EvaluatedArgument() = %q, want %q", got, "before")
	}
}

func TestRecoverAtBoundary(t *testing.T) {
	err := control.RunSafely(func() {
		control.RequirePositive(-1)
	})

	var panicErr *control.PanicError
	matched := errors.As(err, &panicErr)
	t.Logf("error=%q matched=%t", err, matched)

	if !matched {
		t.Fatal("RunSafely() should return *PanicError")
	}
	t.Logf("panic value=%v", panicErr.Value)
	if panicErr.Value != "negative value" {
		t.Errorf("panic value = %v, want %q", panicErr.Value, "negative value")
	}
}
