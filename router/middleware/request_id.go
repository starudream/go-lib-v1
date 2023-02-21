package middleware

import (
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/seq"
)

func RequestId(c *Context) {
	rid := c.GetHeader("X-Request-ID")
	if rid == "" {
		rid = seq.UUID()
		c.Request.Header.Set("X-Request-ID", rid)
	}

	c.Header("X-Request-ID", rid)

	c.Request = c.Request.WithContext(log.With().Str("rid", rid).Logger().WithContext(c))

	c.Next()
}
