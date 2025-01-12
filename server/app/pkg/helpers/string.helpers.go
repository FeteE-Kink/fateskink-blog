package helpers

import (
	"strings"
	"unicode"
)

func IsCamelCase(s string) bool {
	if s == "" {
		return false
	}
	if !unicode.IsLower(rune(s[0])) {
		return false
	}

	for _, r := range s {
		if r == ' ' || r == '_' {
			return false
		}
	}
	return true
}
func CamelToPascalCase(s string) string {
	if !IsCamelCase(s) {
		return s
	}

	return strings.ToUpper(string(s[0])) + s[1:]
}
