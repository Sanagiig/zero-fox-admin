package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zero-fox-admin/api/admin/internal/config"
	"zero-fox-admin/api/admin/internal/middleware"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,
	}
}
