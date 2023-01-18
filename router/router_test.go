package router

import (
	"net/http"
	"testing"

	"github.com/starudream/go-lib/testx"
)

func Test(t *testing.T) {
	Handle(http.MethodGet, "/", func(c *Context) { c.JOK(M{"foo": "bar"}) })
	Handle(http.MethodGet, "/hello", func(c *Context) { c.JOK(M{"bar": "foo"}) })

	T(t,
		TCase{
			Method: http.MethodGet,
			Path:   "/",
			Dump:   true,
			Verify: func(t *testing.T, resp *http.Response, code int, body string) {
				testx.AssertEqualf(t, http.StatusOK, code, "status code")
			},
		},
		TCase{
			Method: http.MethodGet,
			Path:   "/hello",
			Dump:   true,
		},
	)
}
