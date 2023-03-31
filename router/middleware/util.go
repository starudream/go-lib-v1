package middleware

import (
	"regexp"
	"unicode"

	"github.com/starudream/go-lib/internal/gin"
)

type Context = gin.Context

func filterFlags(content string) string {
	for i, char := range content {
		if char == ' ' || char == ';' {
			return content[:i]
		}
	}
	return content
}

func isASCII(s []byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

var jsonCheck = regexp.MustCompile(`(?i:(application|text)/(json|.*\+json|json-.*)(;|$))`)

func isJSONType(ct string) bool {
	return jsonCheck.MatchString(ct)
}
