package router

import (
	"net/http"
	p "path"
)

type Group struct {
	prefix      string
	middlewares []func(http.Handler) http.Handler
}

func NewGroup(prefix string, middlewares ...func(http.Handler) http.Handler) *Group {
	return &Group{
		prefix:      prefix,
		middlewares: middlewares,
	}
}

func (g *Group) Handle(method, pattern string, handle Handler) {
	_r.With(g.middlewares...).MethodFunc(method, p.Join(g.prefix, pattern), wrapHandle(handle))
}
