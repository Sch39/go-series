package main_test

import (
	"testing"
	// "github.com/sch39/go-series/g0101/textkit/internal/normalize"
)

func TestCannotImportInternal(t *testing.T) {
	// _ = normalize.String
	// Compiler menolak import path yang mengandung internal/
	// Test ini memverifikasi: jika baris import di-uncomment, kompilasi gagal.
}
