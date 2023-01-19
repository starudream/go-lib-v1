package router

import (
	"net/http"
	"testing"
)

func TestGroup(t *testing.T) {
	g1 := NewGroup("/v1", Logger)
	g1.Handle(http.MethodGet, "/hello", func(c *Context) { c.TEXT("world") })
	g2 := NewGroup("/v2")
	g2.Handle(http.MethodPost, "/ping", func(c *Context) { c.JSONOK(M{"foo": "bar"}) })

	T(t,
		TCase{
			Method: http.MethodGet,
			Path:   "/v1/hello",
			Dump:   true,
		},
		TCase{
			Method: http.MethodPost,
			Path:   "/v2/ping",
			Dump:   true,
		},
	)
}
