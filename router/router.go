package router

import (
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"

	"github.com/starudream/go-lib/log"
)

type (
	Param  = httprouter.Param
	Params = httprouter.Params
)

type Handler func(c *Context)

type Middleware func(handle Handler) Handler

var _r *httprouter.Router

func init() {
	_r = httprouter.New()

	_r.HandleOPTIONS = true
	_r.HandleMethodNotAllowed = true

	_r.GlobalOPTIONS = handleOPTIONS()
	_r.NotFound = handleNotFound()
	_r.MethodNotAllowed = handleNotAllowed()
	_r.PanicHandler = handlePanic()
}

func R() *httprouter.Router {
	return _r
}

func Handle(method, path string, handle Handler) {
	_r.Handle(method, path, wrapHandle(handle))
}

func handleOPTIONS() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATH, DELETE, HEAD, OPTIONS")
			w.Header().Set("Access-Control-Max-Age", "43200")
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func handleNotFound() http.Handler {
	return wrapHandler(func(c *Context) {
		c.JSON(http.StatusNotFound, ErrNotFound)
	})
}

func handleNotAllowed() http.Handler {
	return wrapHandler(func(c *Context) {
		c.JSON(http.StatusMethodNotAllowed, ErrMethodNotAllowed)
	})
}

func handlePanic() func(w http.ResponseWriter, r *http.Request, rcv any) {
	return func(w http.ResponseWriter, r *http.Request, rcv any) {
		log.Error().Msgf("panic: %s", debug.Stack())
		(&Context{Writer: w}).JSON(http.StatusInternalServerError, ErrInternal)
	}
}

func wrapHandle(handle Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps Params) {
		c := &Context{
			Request: r,
			Writer:  w,
			values:  map[string]any{},
			params:  ps,
		}
		handle(c)
	}
}

func wrapHandler(handle Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := &Context{
			Request: r,
			Writer:  w,
			values:  map[string]any{},
		}
		handle(c)
	})
}
