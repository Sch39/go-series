package textkit

import (
	"strings"
	"unicode"

	"github.com/sch39/go-series/g0101/textkit/internal/normalize"
)

func isWordRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func Slugify(s string) string {
	s = normalize.String(s)

	var result strings.Builder
	lastHyphen := false

	for _, r := range s {
		if isWordRune(r) {
			result.WriteRune(r)
			lastHyphen = false
		} else if !lastHyphen && result.Len() > 0 {
			result.WriteByte('-')
			lastHyphen = true
		}
	}
	return strings.TrimRight(result.String(), "-")
}
