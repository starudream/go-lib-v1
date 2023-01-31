package mx

import (
	"net/http"
	"runtime/debug"

	"github.com/starudream/go-lib/log"

	"github.com/starudream/go-lib/router/cx"
	"github.com/starudream/go-lib/router/ex"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := cx.New(r.Context(), r, w)
		defer func() {
			if err := recover(); err != nil {
				log.Ctx(c).Error().Msgf("%s\n%s", err, debug.Stack())
				c.JSON(http.StatusInternalServerError, ex.Internal)
			}
		}()
		next.ServeHTTP(w, c.Request)
	})
}
