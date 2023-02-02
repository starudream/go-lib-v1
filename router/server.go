package router

import (
	"net/http"

	"github.com/starudream/go-lib/errx"
	"github.com/starudream/go-lib/router/middleware"
	"github.com/starudream/go-lib/validator"

	"github.com/starudream/go-lib/internal/gin"
	"github.com/starudream/go-lib/internal/gin/binding"
)

type (
	Context = gin.Context

	HandlerFunc func(c *Context)
)

var _e *gin.Engine

func init() {
	binding.DefaultValidator = validator.V()

	gin.BindErrorHandler = func(c *gin.Context, err error) {
		if es, ok := err.(validator.ValidationErrors); ok {
			var kvs []any
			for _, e := range es {
				kvs = append(kvs, e.Field(), e.Translate(validator.T()))
			}
			c.Error(errx.ErrParam.AppendMetadata(kvs...))
		} else {
			c.Error(errx.ErrParam.WithMessage(err.Error()))
		}
	}

	_e = gin.New()

	_e.ContextWithFallback = true

	_e.Use(
		middleware.RealIP,
		middleware.RequestId,
		middleware.CORS,
		middleware.Logger,
		middleware.Recover,
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
