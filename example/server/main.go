package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/errx"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/router"
)

const addr = ":80"

func main() {
	Register()

	Server(5 * time.Second)

	Get("/")
	Get("/not")
	Get("/v1/hello")
	Get("/v1/admin/verify")
	Get("/v1/foo")
	Get("/v2/health")
	Get("/v2/panic")
}

func Register() {
	router.Handle(http.MethodGet, "/", func(c *router.Context) {
		c.JSON(http.StatusOK, map[string]any{"version": constant.VERSION})
	})

	g1 := router.Group("/v1")
	{
		g1.Handle(http.MethodGet, "/hello", func(c *router.Context) {
			log.Ctx(c).Info().Msg("world")
		})

		g1a := g1.Group("/admin")
		{
			g1a.Handle(http.MethodGet, "/verify", func(c *router.Context) {
				c.Error(errx.ErrUnAuth)
			})
		}

		g1.Handle(http.MethodGet, "/:name", func(c *router.Context) {
			c.OK(map[string]any{"name": c.Param("name")})
		})
	}

	g2 := router.Group("/v2")
	{
		g2.Handle(http.MethodGet, "/health", func(c *router.Context) {
			c.String(http.StatusOK, c.FullPath())
		})

		g2.Handle(http.MethodGet, "/panic", func(c *router.Context) {
			panic("test")
		})
	}
}

func Server(timeout time.Duration) {
	s := &http.Server{Addr: addr, Handler: router.Handler()}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		err := s.Serve(ln)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}

		time.Sleep(timeout)

		err = s.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}()
}

func Get(path string) {
	resp, err := http.Get("http://localhost" + addr + path)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bs, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Println(">>>\n" + string(bs) + "\n<<<")
}
