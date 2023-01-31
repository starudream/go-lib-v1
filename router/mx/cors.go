package mx

import (
	"net/http"

	"github.com/go-chi/cors"
)

var corsOptions = cors.Options{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{
		http.MethodHead,
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	},
	AllowedHeaders:   []string{"*"},
	AllowCredentials: true,
	MaxAge:           12 * 60 * 60,
}

func CORS(next http.Handler) http.Handler {
	h := cors.Handler(corsOptions)(next)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
