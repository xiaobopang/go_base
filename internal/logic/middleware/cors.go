package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

func (s *sMiddleware) Cors(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{}
	if !r.Response.CORSAllowedOrigin(corsOptions) {
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
