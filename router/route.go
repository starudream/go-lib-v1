package router

import (
	"github.com/starudream/go-lib/internal/gin"
)

type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) IRouter
}

type Router struct {
	x *gin.RouterGroup
}

var _ IRouter = (*Router)(nil)

func (r *Router) Use(middlewares ...HandlerFunc) IRoutes {
	return &Routes{r.x.Use(wrapHF(middlewares)...)}
}

func (r *Router) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	r.x.Handle(httpMethod, relativePath, wrapHF(handlers)...)
}

func (r *Router) Group(relativePath string, handlers ...HandlerFunc) IRouter {
	return &Router{r.x.Group(relativePath, wrapHF(handlers)...)}
}

type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc)
}

type Routes struct {
	x gin.IRoutes
}

var _ IRoutes = (*Routes)(nil)

func (r *Routes) Use(middlewares ...HandlerFunc) IRoutes {
	return &Routes{r.x.Use(wrapHF(middlewares)...)}
}

func (r *Routes) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	r.x.Handle(httpMethod, relativePath, wrapHF(handlers)...)
}
