package middleware

import "go-base/internal/service"

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(new())
}

func new() *sMiddleware {
	return &sMiddleware{}
}
