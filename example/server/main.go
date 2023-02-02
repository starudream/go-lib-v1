package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/errx"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/router"
)

const addr = ":80"

func main() {
	Register()
	Server()

	Do(http.MethodGet, "/")
	Do(http.MethodGet, "/not")
	Do(http.MethodGet, "/v1/log")
	Do(http.MethodGet, "/v1/admin/verify")
	Do(http.MethodGet, "/v1/foo")
	Do(http.MethodGet, "/v2/health")
	Do(http.MethodGet, "/v2/panic")

	Do(http.MethodPost, "/v2/validate", `{"nick": "jack"}`)

}

type User struct {
	Name string `json:"name" validate:"required"`
}

func Register() {
	router.Handle(http.MethodGet, "/", func(c *router.Context) {
		c.JSON(http.StatusOK, map[string]any{"version": constant.VERSION})
	})

	g1 := router.Group("/v1")
	{
		g1.Handle(http.MethodGet, "/log", func(c *router.Context) {
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

		g2.Handle(http.MethodPost, "/validate", func(c *router.Context) {
			if c.BindJSON(&User{}) != nil {
				return
			}
		})
	}
}

func Server() {
	s := &http.Server{Addr: addr, Handler: router.Handler()}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		err = s.Serve(ln)
		if err != nil {
			panic(err)
		}
	}()
}

var client = &http.Client{Timeout: 10 * time.Second}

func Do(method, path string, body ...any) {
	var bodyReader io.Reader
	if len(body) > 0 {
		switch v := body[0].(type) {
		case string:
			bodyReader = strings.NewReader(v)
		case []byte:
			bodyReader = bytes.NewReader(v)
		}
	}

	req, err := http.NewRequest(method, "http://localhost"+addr+path, bodyReader)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
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
