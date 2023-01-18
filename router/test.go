package router

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
	"testing"
)

type TCase struct {
	Method string
	Path   string
	Body   string

	Dump bool

	WarpReq func(t *testing.T, req *http.Request)
	Verify  func(t *testing.T, resp *http.Response, code int, body string)
}

func T(t *testing.T, cases ...TCase) {
	t.Helper()

	for i := 0; i < len(cases); i++ {
		c := cases[i]

		req, err := http.NewRequest(c.Method, c.Path, strings.NewReader(c.Body))
		if err != nil {
			t.Fatal(err)
		}

		if c.WarpReq != nil {
			c.WarpReq(t, req)
		}

		resp := httptest.NewRecorder()

		R().ServeHTTP(resp, req)

		if c.Dump {
			bs1, qe := httputil.DumpRequest(req, true)
			if qe != nil {
				t.Fatal(err)
			}
			t.Log("request:\n" + string(bs1))
			bs2, pe := httputil.DumpResponse(resp.Result(), true)
			if pe != nil {
				t.Fatal(err)
			}
			t.Log("response:\n" + string(bs2))
		}

		if c.Verify != nil {
			c.Verify(t, resp.Result(), resp.Code, resp.Body.String())
		}
	}
}
