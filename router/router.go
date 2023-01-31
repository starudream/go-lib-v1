package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/starudream/go-lib/router/cx"
	"github.com/starudream/go-lib/router/ex"
	"github.com/starudream/go-lib/router/mx"
)

type Handler func(c *cx.Context)

var _r *chi.Mux

func init() {
	_r = chi.NewRouter()

	_r.Use(
		middleware.RealIP,
		mx.RequestId,
		mx.Recover,
		mx.CORS,
		mx.Logger,
	)

	_r.NotFound(handleNotFound())
	_r.MethodNotAllowed(handleMethodNotAllowed())
}

func R() *chi.Mux {
	return _r
}

func handleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cx.FromRequest(r).JSON(http.StatusNotFound, ex.NotFound)
	}
}

func handleMethodNotAllowed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cx.FromRequest(r).JSON(http.StatusMethodNotAllowed, ex.MethodNotAllowed)
	}
}

func Handle(method, pattern string, handle Handler) {
	_r.MethodFunc(method, pattern, wrapHandle(handle))
}

func wrapHandle(handle Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handle(cx.FromRequest(r))
	}
}
