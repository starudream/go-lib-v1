package router

import (
	"net/http"
	"testing"

	"github.com/starudream/go-lib/testx"
)

func Test(t *testing.T) {
	Handle(http.MethodGet, "/", func(c *Context) { c.JSONOK(M{"foo": "bar"}) })
	Handle(http.MethodGet, "/hello", func(c *Context) { c.JSONOK(M{"bar": "foo"}) })
	Handle(http.MethodGet, "/download", func(c *Context) { c.FILE("./util.go") })
	Handle(http.MethodGet, "/attachment", func(c *Context) { c.ATTACHMENT("./util.go", "util.go") })

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
		TCase{
			Method: http.MethodGet,
			Path:   "/download",
			Dump:   true,
		},
		TCase{
			Method: http.MethodGet,
			Path:   "/attachment",
			Dump:   true,
		},
	)
}
