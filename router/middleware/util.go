package middleware

import (
	"net/http"
	"unicode"

	"github.com/starudream/go-lib/internal/gin"
)

type Context = gin.Context

var (
	contentType   = http.CanonicalHeaderKey("Content-Type")
	trueClientIP  = http.CanonicalHeaderKey("True-Client-IP")
	xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
	xRealIP       = http.CanonicalHeaderKey("X-Real-IP")
	xRequestID    = http.CanonicalHeaderKey("X-Request-ID")
)

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
