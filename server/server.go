package server

import (
	"net/http"

	"github.com/starudream/go-lib/errx"
	"github.com/starudream/go-lib/server/middleware"

	"github.com/starudream/go-lib/internal/gin"
)

type (
	Context = gin.Context

	HandlerFunc func(c *Context)
)

var _e *gin.Engine

func init() {
	_e = gin.New()

	_e.UseH2C = true
	_e.ContextWithFallback = true

	_e.Use(
		middleware.RealIP,
		middleware.RequestId,
		middleware.Recover,
		middleware.CORS,
		middleware.Logger,
	)

	_e.NoRoute(func(c *Context) { c.Error(errx.ErrNotFound) })
	_e.NoMethod(func(c *Context) { c.Error(errx.ErrNoMethod) })
}

func E() *gin.Engine {
	return _e
}

func Handler() http.Handler {
	return _e.Handler()
}

func Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	_e.Handle(httpMethod, relativePath, wrapHF(handlers)...)
}

func Group(relativePath string, handlers ...HandlerFunc) IRouter {
	return &Router{_e.Group(relativePath, wrapHF(handlers)...)}
}

func Use(middlewares ...HandlerFunc) IRoutes {
	return &Routes{_e.Use(wrapHF(middlewares)...)}
}
