package textkit_test

import (
	"testing"

	"github.com/sch39/go-series/g0101/textkit"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"Hello World!", "hello-world"},
		{"FOO---BAR", "foo-bar"},
		{"  trim  ", "trim"},
	}
	for _, tt := range tests {
		if got := textkit.Slugify(tt.input); got != tt.want {
			t.Errorf("Slugify(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
