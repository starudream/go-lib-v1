package mx

import (
	"net/http"
	"strings"
	"time"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/log"

	"github.com/starudream/go-lib/router/cx"
	"github.com/starudream/go-lib/router/ex"

	"github.com/starudream/go-lib/internal/httpresponse"
)

func Logger(next http.Handler) http.Handler {
	ml := config.GetInt("router.log.resp_max_lines")
	if ml <= 0 {
		ml = 16
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := cx.FromRequest(r)

		start := time.Now()

		bs, err := c.BodyBytes()
		if err != nil {
			log.Error().Msgf("read http body error: %v", err)
			c.JSON(http.StatusInternalServerError, ex.Internal)
			return
		}

		l := log.
			Ctx(c).
			With().
			Str("method", r.Method).
			Str("path", r.URL.String()).
			Str("remote", r.RemoteAddr).
			Logger()

		if len(bs) == 0 {
			bs = []byte("<empty>")
		} else if req, ue := json.UnmarshalTo[any](bs); ue == nil {
			bs = json.MustMarshal(req)
		}

		l.Info().Msgf("req=%s", bs)

		c.Writer = httpresponse.NewResponse(c.Writer)

		next.ServeHTTP(c.Writer, c.Request)

		resp, sc := httpresponse.GetResponse(c.Writer)

		if len(resp) == 0 {
			resp = []byte("<empty>")
		} else if strings.Count(string(resp), "\n") > ml {
			resp = []byte("<ignore>")
		}

		l.Info().Int("code", sc).Dur("took", time.Since(start)).Msgf("resp=%s", resp)
	})
}
