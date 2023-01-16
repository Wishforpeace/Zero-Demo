package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type ExampleMiddleware struct {
}

func NewExampleMiddleware() *ExampleMiddleware {
	return &ExampleMiddleware{}
}

func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Infov("example middle")
		// Passthrough to next handler if need
		next(w, r)
	}
}