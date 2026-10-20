package faults_test

import (
	"testing"

	"github.com/sch39/go-series/g0106/errors/faults"
)

func TestSentinelAndWrappedError(t *testing.T) {
	directErr := faults.FindUser(7)
	wrappedErr := faults.LoadProfile(7)
	directMatch := faults.IsNotFound(directErr)
	wrappedMatch := faults.IsNotFound(wrappedErr)

	t.Logf("direct=%q match=%t", directErr, directMatch)
	t.Logf("wrapped=%q match=%t", wrappedErr, wrappedMatch)

	if !directMatch {
		t.Error("FindUser() error should match ErrNotFound")
	}
	if !wrappedMatch {
		t.Error("LoadProfile() wrapped error should match ErrNotFound")
	}
}

func TestTypedErrorWithErrorsAs(t *testing.T) {
	err := faults.Register("")
	validationErr, ok := faults.ValidationDetails(err)

	t.Logf("error=%q matched=%t", err, ok)
	if ok {
		t.Logf("field=%q value=%q", validationErr.Field, validationErr.Value)
	}

	if !ok {
		t.Fatal("Register() error should contain *ValidationError")
	}
	if validationErr.Field != "email" || validationErr.Value != "" {
		t.Errorf("ValidationError = %+v, want email with empty value", validationErr)
	}
}

func TestSuccessReturnsNil(t *testing.T) {
	findErr := faults.FindUser(42)
	registerErr := faults.Register("dev@example.com")

	t.Logf("findErr=%v registerErr=%v", findErr, registerErr)

	if findErr != nil || registerErr != nil {
		t.Errorf("success should return nil: find=%v register=%v", findErr, registerErr)
	}
}
