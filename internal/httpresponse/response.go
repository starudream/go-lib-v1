package httpresponse

import (
	"net/http"
)

type Response struct {
	w http.ResponseWriter

	bs []byte
	sc int
}

var _ http.ResponseWriter = (*Response)(nil)

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w: w}
}

func GetResponse(w http.ResponseWriter) ([]byte, int) {
	v, ok := w.(*Response)
	if ok {
		return v.bs, v.sc
	}
	return nil, 0
}

func (w *Response) Header() http.Header {
	return w.w.Header()
}

func (w *Response) Write(bs []byte) (int, error) {
	w.bs = bs
	return w.w.Write(bs)
}

func (w *Response) WriteHeader(statusCode int) {
	w.sc = statusCode
	w.w.WriteHeader(statusCode)
}
