package middleware

import (
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/seq"
)

func RequestId(c *Context) {
	rid := c.GetHeader(xRequestID)
	if rid == "" {
		rid = seq.UUID()
		c.Request.Header.Set(xRequestID, rid)
	}
	c.Header(xRequestID, rid)
	c.Request = c.Request.WithContext(log.With().Str("rid", rid).Logger().WithContext(c))

	c.Next()
}
