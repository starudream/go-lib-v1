package errx

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`

	Metadata map[string]any `json:"metadata,omitempty"`

	statusCode int
}

func New(code int, s string, v ...any) *Error {
	return &Error{Code: code, Message: format(s, v...)}
}

func From(err error) *Error {
	if err == nil {
		return nil
	}
	if e := new(Error); errors.As(err, &e) {
		return e
	}
	return ErrInternal.WithMessage(err.Error())
}

var _ error = (*Error)(nil)

func (e *Error) Error() (s string) {
	s = "code: " + strconv.Itoa(e.Code)
	if e.Message != "" {
		s += ", message: " + e.Message
	}
	if len(e.Metadata) > 0 {
		s += ", metadata:"
		for k, v := range e.Metadata {
			s += " " + k + "=" + fmt.Sprintf("%v", v)
		}
	}
	return
}

func (e *Error) Copy() *Error {
	ne := new(Error)
	*ne = *e
	if ne.Metadata == nil {
		ne.Metadata = map[string]any{}
	}
	return ne
}

func (e *Error) WithStatusCode(statusCode int) *Error {
	ne := e.Copy()
	ne.statusCode = statusCode
	return ne
}

func (e *Error) WithMessage(s string, v ...any) *Error {
	ne := e.Copy()
	ne.Message = format(s, v...)
	return ne
}

func (e *Error) WithMetadata(kvs ...any) *Error {
	ne := e.Copy()
	ne.Metadata = map[string]any{}
	return ne.AppendMetadata(kvs...)
}

func (e *Error) AppendMetadata(kvs ...any) *Error {
	if len(kvs)%2 != 0 {
		panic("kvs must be even")
	}
	ne := e.Copy()
	for i := 0; i < len(kvs); i += 2 {
		ki, vi := kvs[i], kvs[i+1]
		k, ok := ki.(string)
		if !ok {
			k = fmt.Sprint(k)
		}
		ne.Metadata[k] = vi
	}
	return ne
}

func (e *Error) StatusCode() int {
	if e == nil || e.statusCode == 0 {
		return http.StatusInternalServerError
	}
	return e.statusCode
}

func format(s string, v ...any) string {
	if len(v) == 0 {
		return s
	}
	return fmt.Sprintf(s, v...)
}
