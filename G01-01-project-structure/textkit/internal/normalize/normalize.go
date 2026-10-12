package normalize

import "strings"

func String(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
