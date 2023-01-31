package router

import (
	"net/http"
	"testing"

	"github.com/starudream/go-lib/testx"

	"github.com/starudream/go-lib/router/cx"
	"github.com/starudream/go-lib/router/dx"
)

func Test(t *testing.T) {
	Handle(http.MethodGet, "/", func(c *cx.Context) { c.JSONOK(dx.M{"foo": "bar"}) })
	Handle(http.MethodGet, "/hello", func(c *cx.Context) { c.JSONOK(dx.M{"bar": "foo"}) })
	Handle(http.MethodGet, "/{args}", func(c *cx.Context) { c.JSONOK(dx.M{"args": c.Param("args")}) })
	Handle(http.MethodGet, "/download", func(c *cx.Context) { c.FILE("./test.go") })
	Handle(http.MethodGet, "/attachment", func(c *cx.Context) { c.ATTACHMENT("./test.go", "test.go") })

	TE(t,
		TC{
			Method:  http.MethodGet,
			Pattern: "/",
			Dump:    true,
			Verify: func(t *testing.T, resp *http.Response, code int, body string) {
				testx.AssertEqualf(t, http.StatusOK, code, "status code")
			},
		},
		TC{
			Method:  http.MethodGet,
			Pattern: "/hello",
		},
		TC{
			Method:  http.MethodGet,
			Pattern: "/version",
		},
		TC{
			Method:  http.MethodGet,
			Pattern: "/download",
		},
		TC{
			Method:  http.MethodGet,
			Pattern: "/attachment",
		},
	)
}
