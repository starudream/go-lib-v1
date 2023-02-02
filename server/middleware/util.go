package middleware

import (
	"net/http"

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
