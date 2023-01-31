package mx

import (
	"net/http"

	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/seq"

	"github.com/starudream/go-lib/router/dx"
)

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tid := r.Header.Get(dx.XRequestId)
		if tid == "" {
			tid = seq.UUID()
			w.Header().Set(dx.XRequestId, tid)
		}
		next.ServeHTTP(w, r.WithContext(log.With().Str("tid", tid).Logger().WithContext(r.Context())))
	})
}
