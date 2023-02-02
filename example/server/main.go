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
	"github.com/starudream/go-lib/server"
)

const addr = ":80"

func main() {
	Register()

	Server(5 * time.Second)

	Get("/")
	Get("/not")
	Get("/v1/hello")
	Get("/v1/admin/verify")
	Get("/v2/health")
}

func Register() {
	server.Handle(http.MethodGet, "/", func(c *server.Context) {
		c.JSON(http.StatusOK, map[string]any{"version": constant.VERSION})
	})

	g1 := server.Group("/v1")
	{
		g1.Handle(http.MethodGet, "/hello", func(c *server.Context) {
			log.Ctx(c).Info().Msg("world")
			c.Error(errx.OK)
		})

		g1a := g1.Group("/admin")
		{
			g1a.Handle(http.MethodGet, "/verify", func(c *server.Context) {
				c.Error(errx.ErrUnAuth)
			})
		}
	}

	g2 := server.Group("/v2")
	{
		g2.Handle(http.MethodGet, "/health", func(c *server.Context) {
			c.String(http.StatusOK, "ok")
		})
	}
}

func Server(timeout time.Duration) {
	s := &http.Server{Addr: addr, Handler: server.Handler()}

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
