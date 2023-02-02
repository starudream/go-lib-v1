package gin

import (
	"bytes"
	"net/http"

	"github.com/starudream/go-lib/log"
)

type ResponseWriter interface {
	http.ResponseWriter

	// Size returns the number of bytes already written into the response http body.
	Size() int

	// Data returns the response body already written into the response http body.
	Data() []byte

	// Status returns the HTTP response status code of the current request.
	Status() int

	// Written returns true if the response body was already written.
	Written() bool

	// WriteHeaderNow forces to write the http header (status code + headers).
	WriteHeaderNow()
}

type responseWriter struct {
	w http.ResponseWriter

	size   int
	status int
	buffer bytes.Buffer
}

var _ ResponseWriter = (*responseWriter)(nil)

func (w *responseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *responseWriter) Write(bs []byte) (n int, err error) {
	w.WriteHeaderNow()
	n, err = w.w.Write(bs)
	w.size += n
	w.buffer.Write(bs)
	return
}

func (w *responseWriter) WriteHeader(code int) {
	if code > 0 && w.status != code {
		if w.Written() {
			log.Warn().Msgf("headers were already written, wanted to override status code %d with %d", w.status, code)
		}
		w.status = code
	}
}

func (w *responseWriter) Size() int {
	return w.size
}

func (w *responseWriter) Data() []byte {
	return w.buffer.Bytes()
}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) Written() bool {
	return w.size != -1
}

func (w *responseWriter) WriteHeaderNow() {
	if !w.Written() {
		w.size = 0
		w.w.WriteHeader(w.status)
	}
}

func (w *responseWriter) reset(nw http.ResponseWriter) {
	w.w = nw
	w.size = -1
	w.status = http.StatusOK
	w.buffer.Reset()
}
