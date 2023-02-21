package middleware

import (
	"net"
	"net/http"
	"strings"
)

func RealIP(c *Context) {
	if rip := realIP(c.Request); rip != "" {
		c.Request.RemoteAddr = rip
	}
	c.Next()
}

func realIP(r *http.Request) string {
	var ip string

	if tcip := r.Header.Get("True-Client-IP"); tcip != "" {
		ip = tcip
	} else if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		ip = xrip
	} else if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		i := strings.Index(xff, ",")
		if i == -1 {
			i = len(xff)
		}
		ip = xff[:i]
	}

	if ip == "" || net.ParseIP(ip) == nil {
		return ""
	}

	return ip
}
