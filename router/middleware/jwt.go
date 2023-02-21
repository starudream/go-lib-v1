package middleware

import (
	"github.com/starudream/go-lib/errx"
	"github.com/starudream/go-lib/jwt"
	"github.com/starudream/go-lib/log"
)

func JWT(c *Context) {
	raw := c.GetHeader("Authorization")
	if raw == "" {
		raw = c.GetHeader("Token")
	}

	if raw == "" {
		c.AbortWithError(errx.ErrUnAuth.WithMessage("missing token"))
		return
	}

	claims, err := jwt.Parse(raw)
	if err != nil {
		c.AbortWithError(errx.ErrUnAuth.WithMessage("invalid token"))
		return
	}

	l := log.Ctx(c).With().Str("iss", claims.Issuer).Str("sub", claims.Subject).Logger()

	c.Request = c.Request.WithContext(l.WithContext(claims.WithContext(c)))

	c.Next()
}
