package errx

import (
	"net/http"
)

const statusCodeBegin = 1000

var (
	OK = New(0, "ok").WithStatusCode(http.StatusOK)

	ErrInternal = New(statusCodeBegin, "Internal Error").WithStatusCode(http.StatusInternalServerError)

	ErrParam     *Error
	ErrUnAuth    *Error
	ErrForbidden *Error
	ErrNotFound  *Error
	ErrNoMethod  *Error
	ErrConflict  *Error
)

func init() {
	ErrParam = FromStatusCode(http.StatusBadRequest)
	ErrUnAuth = FromStatusCode(http.StatusUnauthorized)
	ErrForbidden = FromStatusCode(http.StatusForbidden)
	ErrNotFound = FromStatusCode(http.StatusNotFound)
	ErrNoMethod = FromStatusCode(http.StatusMethodNotAllowed)
	ErrConflict = FromStatusCode(http.StatusConflict)
}

func FromStatusCode(statusCode int) *Error {
	text := http.StatusText(statusCode)
	if text == "" {
		return ErrInternal
	}
	return New(statusCodeBegin+statusCode, text).WithStatusCode(statusCode)
}
