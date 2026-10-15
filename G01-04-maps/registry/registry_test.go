package registry_test

import (
	"testing"

	"github.com/sch39/go-series/g0104/registry"
)

func TestLookupDistinguishesMissingFromZero(t *testing.T) {
	r := registry.New()
	r.Put("zero", 0)

	value, ok := r.Lookup("zero")
	if !ok || value != 0 {
		t.Fatalf("Lookup(zero) = %d, %v; want 0, true", value, ok)
	}

	missing, present := r.Lookup("missing")
	t.Logf("zero=(%d,%v) missing=(%d,%v)", value, ok, missing, present)
	if present || missing != 0 {
		t.Errorf("Lookup(missing) = %d, %v; want 0, false", missing, present)
	}
}

func TestDeleteAndConcurrentAccess(t *testing.T) {
	r := registry.New()
	r.Put("temporary", 7)
	r.Delete("temporary")
	if _, ok := r.Lookup("temporary"); ok {
		t.Fatal("Delete did not remove temporary")
	}

	done := make(chan struct{})
	for i := 0; i < 4; i++ {
		go func(i int) {
			for j := 0; j < 100; j++ {
				r.Put("shared", i+j)
				_, _ = r.Lookup("shared")
			}
			done <- struct{}{}
		}(i)
	}
	for i := 0; i < 4; i++ {
		<-done
	}
	value, ok := r.Lookup("shared")
	t.Logf("shared=(%d,%v) after concurrent access", value, ok)
	if !ok {
		t.Error("shared key missing after concurrent access")
	}
}
