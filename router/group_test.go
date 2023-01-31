package router

import (
	"net/http"
	"testing"

	"github.com/starudream/go-lib/log"

	"github.com/starudream/go-lib/router/cx"
	"github.com/starudream/go-lib/router/dx"
)

func TestGroup(t *testing.T) {
	g1 := NewGroup("/v1")
	g1.Handle(http.MethodGet, "/hello", func(c *cx.Context) { c.TEXT("world") })
	g1.Handle(http.MethodGet, "/tid", func(c *cx.Context) { log.Ctx(c).Info().Msgf("tid") })
	g2 := NewGroup("/v2")
	g2.Handle(http.MethodPost, "/ping", func(c *cx.Context) { c.JSONOK(dx.M{"foo": "bar"}) })

	TE(t,
		TC{
			Method:  http.MethodGet,
			Pattern: "/v1/hello",
		},
		TC{
			Method:  http.MethodGet,
			Pattern: "/v1/tid",
		},
		TC{
			Method:  http.MethodPost,
			Pattern: "/v2/ping",
			Dump:    true,
		},
	)
}
