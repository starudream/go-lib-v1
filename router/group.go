package router

import (
	p "path"
)

type Group struct {
	prefix      string
	middlewares []Middleware
}

func NewGroup(prefix string, middlewares ...Middleware) *Group {
	if prefix == "" {
		prefix = "/"
	}
	return &Group{
		prefix:      prefix,
		middlewares: middlewares,
	}
}

func (g *Group) Handle(method, path string, handle Handler) {
	if ml := len(g.middlewares); ml > 0 {
		for i := ml - 1; i >= 0; i-- {
			handle = g.middlewares[i](handle)
		}
	}
	R().Handle(method, p.Join(g.prefix, path), wrapHandle(handle))
}
