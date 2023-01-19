package router

import (
	"net/http"
	"time"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/seq"

	"github.com/starudream/go-lib/internal/httpresponse"
)

func Logger(h Handler) Handler {
	return func(c *Context) {
		start := time.Now()

		bs, err := c.BodyBytes()
		if err != nil {
			log.Error().Msgf("read http body error: %v", err)
			c.JSON(http.StatusInternalServerError, ErrInternal)
			return
		}

		tid := c.GetRequestId()
		if tid == "" {
			tid = seq.UUID()
			c.SetHeader("X-Request-ID", tid)
		}

		r := c.Request

		l := log.
			With().
			Str("method", r.Method).
			Str("path", r.URL.String()).
			Str("remote", r.RemoteAddr).
			Str("type", c.GetContentType()).
			Str("tid", tid).
			Logger()

		var req any

		if json.Unmarshal(bs, &req) == nil {
			l.Info().Msgf("req=%s", json.MustMarshal(req))
		} else {
			l.Info().Msgf("req=%s", bs)
		}

		c.Writer = httpresponse.NewResponse(c.Writer)

		nc := c.WithContext(log.With().Str("tid", tid).Logger().WithContext(c))

		h(nc)

		resp, sc := httpresponse.GetResponse(c.Writer)

		l.Info().Int("code", sc).Dur("took", time.Since(start)).Msgf("resp=%s", resp)
	}
}
