package jwt

import (
	"context"

	"github.com/starudream/go-lib/errx"
)

type claimsKey struct{}

func (c *Claims) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, claimsKey{}, c)
}

func FromContext(ctx context.Context) (*Claims, error) {
	c, ok := ctx.Value(claimsKey{}).(*Claims)
	if ok {
		return c, nil
	}
	return nil, errx.ErrUnAuth.WithMessage("missing token")
}
