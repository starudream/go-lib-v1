package cx

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cast"

	"github.com/starudream/go-lib/codec/json"

	"github.com/starudream/go-lib/router/dx"
)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter

	ctx    *chi.Context
	mu     sync.RWMutex
	values map[string]any
	query  url.Values
}

type ctxKey struct{}

func New(ctx context.Context, r *http.Request, w http.ResponseWriter) *Context {
	if ctx == nil {
		ctx = context.Background()
	}

	if c, ok := ctx.(*Context); ok {
		return c
	}

	if c, ok := ctx.Value(ctxKey{}).(*Context); ok {
		return c
	}

	c := &Context{
		Request: r,
		Writer:  w,
		values:  map[string]any{},
	}

	if r != nil {
		c.ctx = chi.RouteContext(r.Context())
		c.query = r.URL.Query()
	}

	ctx = context.WithValue(ctx, ctxKey{}, c)

	if r != nil {
		c.Request = r.WithContext(ctx)
	}

	return c
}

func FromRequest(req *http.Request) *Context {
	return New(req.Context(), nil, nil)
}

// --------------------------------------------------------------------------------

var _ context.Context = (*Context)(nil)

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	if c.Request == nil {
		return
	}
	return c.Request.Context().Deadline()
}

func (c *Context) Done() <-chan struct{} {
	if c.Request == nil {
		return nil
	}
	return c.Request.Context().Done()
}

func (c *Context) Err() error {
	if c.Request == nil {
		return nil
	}
	return c.Request.Context().Err()
}

func (c *Context) Value(key any) any {
	if vk, ok := key.(string); ok {
		if value, exist := c.Get(vk); exist {
			return value
		}
	}
	if c.Request == nil {
		return nil
	}
	return c.Request.Context().Value(key)
}

func (c *Context) WithContext(ctx context.Context) *Context {
	c.Request = c.Request.WithContext(ctx)
	return c
}

// --------------------------------------------------------------------------------

func (c *Context) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.values == nil {
		c.values = map[string]any{}
	}
	c.values[key] = value
}

func (c *Context) Get(key string) (any, bool) {
	c.mu.RLocker()
	defer c.mu.RUnlock()
	value, exist := c.values[key]
	return value, exist
}

func (c *Context) GetString(key string) string {
	value, _ := c.Get(key)
	return cast.ToString(value)
}

func (c *Context) GetBool(key string) bool {
	value, _ := c.Get(key)
	return cast.ToBool(value)
}

func (c *Context) GetInt(key string) int {
	value, _ := c.Get(key)
	return cast.ToInt(value)
}

func (c *Context) GetFloat64(key string) float64 {
	value, _ := c.Get(key)
	return cast.ToFloat64(value)
}

// --------------------------------------------------------------------------------

func (c *Context) Param(key string) string {
	return c.ctx.URLParam(key)
}

func (c *Context) Query(key string) string {
	return c.query.Get(key)
}

func (c *Context) AllQuery() url.Values {
	return c.query
}

func (c *Context) SetHeader(key, value string) {
	if value == "" {
		c.Writer.Header().Del(key)
		return
	}
	c.Writer.Header().Set(key, value)
}

func (c *Context) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

func (c *Context) SetStatusCode(statusCode int) {
	c.Writer.WriteHeader(statusCode)
}

func (c *Context) BodyBytes() ([]byte, error) {
	bs, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bs))
	return bs, nil
}

// --------------------------------------------------------------------------------

func (c *Context) GetContentType() string {
	return c.GetHeader(dx.ContentType)
}

func (c *Context) GetAuthorization() string {
	return stringNotEmpty(
		c.GetHeader(dx.Authorization),
		c.GetHeader("Token"),
		c.Query("authorization"),
		c.Query("token"),
	)
}

func (c *Context) GetRequestId() string {
	return c.GetHeader(dx.XRequestId)
}

// --------------------------------------------------------------------------------

func (c *Context) TEXT(s string, v ...any) {
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write([]byte(format(s, v...)))
}

func (c *Context) JSON(code int, v any) {
	c.Writer.Header().Set(dx.ContentType, dx.ApplicationJSON)
	c.Writer.WriteHeader(code)
	_, _ = c.Writer.Write(json.MustMarshal(v))
}

func (c *Context) JSONOK(v any) {
	c.JSON(http.StatusOK, v)
}

func (c *Context) FILE(filepath string) {
	http.ServeFile(c.Writer, c.Request, filepath)
}

func (c *Context) ATTACHMENT(filepath, filename string) {
	if isASCII(filename) {
		c.Writer.Header().Set(dx.ContentDisposition, `attachment; filename="`+filename+`"`)
	} else {
		c.Writer.Header().Set(dx.ContentDisposition, `attachment; filename*=UTF-8''`+url.QueryEscape(filename))
	}
	http.ServeFile(c.Writer, c.Request, filepath)
}
