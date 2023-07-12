package middleware

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"time"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/errx"
	"github.com/starudream/go-lib/log"
)

var (
	rawDataEmpty  = []byte("<empty>")
	rawDataIgnore = []byte("<ignore>")
	rawDataBinary = []byte("<binary>")
)

func Logger(c *Context) {
	start := time.Now()

	req, err := c.GetRawData()
	if err != nil {
		log.Ctx(c).Error().Msgf("read http body error: %v", err)
		c.AbortWithError(errx.ErrInternal)
		return
	}

	l := log.Ctx(c).With().
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.String()).
		Str("remote", c.Request.RemoteAddr).
		Logger()

	if len(req) == 0 {
		req = rawDataEmpty
	} else if v, ue := json.UnmarshalTo[any](req); ue == nil {
		req = json.MustMarshal(v)
	}

	l.Info().
		Str("type", filterFlags(c.GetHeader("Content-Type"))).
		Msgf("req=%s", req)

	c.Next()

	resp, statusCode := c.Writer.Data(), c.Writer.Status()

	switch c.Writer.Header().Get("Content-Encoding") {
	case "":
	case "gzip":
		resp = gunzip(resp)
	default:
		resp = rawDataBinary
	}

	if isJSONType(c.Writer.Header().Get("Content-Type")) {
	} else if len(resp) == 0 {
		resp = rawDataEmpty
	} else if bytes.Count(resp, []byte("\n")) >= 50 {
		resp = rawDataIgnore
	} else if !isASCII(resp) {
		resp = rawDataBinary
	}

	lvl := log.InfoLevel
	if statusCode != http.StatusOK {
		lvl = log.ErrorLevel
	}

	l.WithLevel(lvl).
		Int("code", statusCode).
		Str("type", filterFlags(c.Writer.Header().Get("Content-Type"))).
		Dur("took", time.Since(start)).
		Msgf("resp=%s", resp)
}

func gunzip(bs []byte) []byte {
	r, err := gzip.NewReader(bytes.NewBuffer(bs))
	if err != nil {
		return rawDataBinary
	}
	defer r.Close()
	v, err := io.ReadAll(r)
	if err != nil {
		return rawDataBinary
	}
	return v
}
