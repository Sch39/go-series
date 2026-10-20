package client_test

import (
	"testing"
	"time"

	"github.com/sch39/go-series/g0109/functionaloptions/client"
)

func TestNewUsesDefaults(t *testing.T) {
	got, err := client.New()
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}
	t.Logf("endpoint=%s timeout=%s retries=%d", got.Endpoint, got.Timeout, got.Retries)

	if got.Endpoint != "https://api.example.com" || got.Timeout != 5*time.Second || got.Retries != 2 {
		t.Errorf("New defaults = %+v", got)
	}
}

func TestOptionsOverrideInOrder(t *testing.T) {
	got, err := client.New(
		client.WithEndpoint("https://internal.example.com"),
		client.WithTimeout(2*time.Second),
		client.WithRetries(4),
		client.WithRetries(1),
	)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}
	t.Logf("endpoint=%s timeout=%s retries=%d", got.Endpoint, got.Timeout, got.Retries)

	if got.Endpoint != "https://internal.example.com" || got.Timeout != 2*time.Second || got.Retries != 1 {
		t.Errorf("configured client = %+v", got)
	}
}

func TestInvalidOptionStopsConstruction(t *testing.T) {
	got, err := client.New(client.WithTimeout(0))
	t.Logf("client=%v error=%v", got, err)

	if err == nil {
		t.Fatal("New returned nil error, want validation error")
	}
	if got != nil {
		t.Errorf("New returned client %v, want nil", got)
	}
}
