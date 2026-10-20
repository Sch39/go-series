package faults

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("record not found")

func FindUser(id int) error {
	if id != 42 {
		return ErrNotFound
	}
	return nil
}

func LoadProfile(id int) error {
	if err := FindUser(id); err != nil {
		return fmt.Errorf("load profile %d: %w", id, err)
	}
	return nil
}

type ValidationError struct {
	Field string
	Value string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s %q", e.Field, e.Value)
}

func ValidateEmail(email string) error {
	if email == "" {
		return &ValidationError{Field: "email", Value: email}
	}
	return nil
}

func Register(email string) error {
	if err := ValidateEmail(email); err != nil {
		return fmt.Errorf("register user: %w", err)
	}
	return nil
}

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func ValidationDetails(err error) (*ValidationError, bool) {
	var validationErr *ValidationError
	if !errors.As(err, &validationErr) {
		return nil, false
	}
	return validationErr, true
}
