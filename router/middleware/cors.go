package middleware

import (
	"net/http"
	"time"

	"github.com/starudream/go-lib/internal/gin/cors"
)

func CORS(c *Context) {
	m := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	m(c)
}
