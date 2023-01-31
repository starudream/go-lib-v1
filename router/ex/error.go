package ex

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cast"
)

var (
	BadRequest       = New(http.StatusBadRequest, "bad request")
	Unauthorized     = New(http.StatusUnauthorized, "unauthorized")
	Forbidden        = New(http.StatusForbidden, "forbidden")
	NotFound         = New(http.StatusNotFound, "not found")
	MethodNotAllowed = New(http.StatusMethodNotAllowed, "method not allowed")
	Conflict         = New(http.StatusConflict, "conflict")
	Internal         = New(http.StatusInternalServerError, "internal server error")
)

type Error struct {
	Code    int    `json:"code" xml:"code"`
	Message string `json:"message,omitempty" xml:"message,omitempty"`

	Metadata map[string]any `json:"metadata,omitempty" xml:"metadata,omitempty"`

	ks []string
}

var _ error = (*Error)(nil)

func New(code int, s string, v ...any) *Error {
	return &Error{Code: code, Message: format(s, v...)}
}

func (e *Error) Error() (s string) {
	s = "code: " + strconv.Itoa(e.Code)
	if e.Message != "" {
		s += ", message: " + e.Message
	}
	if len(e.Metadata) > 0 {
		s += ", metadata:"
		for i := 0; i < len(e.ks); i++ {
			k := e.ks[i]
			v := e.Metadata[k]
			s += " " + k + "=" + fmt.Sprintf("%v", v)
		}
	}
	return
}

func (e *Error) WithMessage(s string, v ...any) *Error {
	e.Message = format(s, v...)
	return e
}

func (e *Error) WithMetadata(kvs ...any) *Error {
	e.Metadata, e.ks = map[string]any{}, []string{}
	return e.AppendMetadata(kvs...)
}

func (e *Error) AppendMetadata(kvs ...any) *Error {
	if len(kvs)%2 != 0 {
		panic("kvs must be even")
	}
	if e.Metadata == nil {
		e.Metadata = map[string]any{}
	}
	for i := 0; i < len(kvs); i += 2 {
		ki, vi := kvs[i], kvs[i+1]
		k, ok := ki.(string)
		if !ok {
			k = cast.ToString(k)
		}
		e.Metadata[k] = vi
		e.ks = append(e.ks, k)
	}
	return e
}

func format(s string, v ...any) string {
	if len(v) == 0 {
		return s
	}
	return fmt.Sprintf(s, v...)
}
