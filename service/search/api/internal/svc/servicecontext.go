package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/service/search/api/internal/config"
	"go-zero-demo/service/search/api/internal/middleware"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
	}
}
