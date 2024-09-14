package svc

import (
	"zero-fox-admin/api/admin/internal/config"
	"zero-fox-admin/api/admin/internal/middleware"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	Redis    *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.Redis.Address,
	})

	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,
		Redis:    rds,
	}
}
