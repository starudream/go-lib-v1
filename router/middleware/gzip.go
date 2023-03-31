package middleware

import (
	"github.com/starudream/go-lib/internal/gin/gzip"
)

func GZIP(c *Context) {
	m := gzip.Gzip(gzip.DefaultCompression)
	m(c)
}
