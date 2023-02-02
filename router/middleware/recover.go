package middleware

import (
	"runtime/debug"

	"github.com/starudream/go-lib/errx"
	"github.com/starudream/go-lib/log"
)

func Recover(c *Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Ctx(c).Error().Msgf("http panic recover: %v\n%s", err, debug.Stack())
			c.AbortWithError(errx.ErrInternal)
		}
	}()

	c.Next()
}
