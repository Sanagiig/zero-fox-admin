package svc

import (
	"zero-fox-admin/api/admin/internal/config"
	"zero-fox-admin/api/admin/internal/middleware"
	"zero-fox-admin/rpc/sys/client/deptservice"
	"zero-fox-admin/rpc/sys/client/dictitemservice"
	"zero-fox-admin/rpc/sys/client/dicttypeservice"
	"zero-fox-admin/rpc/sys/client/loginlogservice"
	"zero-fox-admin/rpc/sys/client/menuservice"
	"zero-fox-admin/rpc/sys/client/operatelogservice"
	"zero-fox-admin/rpc/sys/client/postservice"
	"zero-fox-admin/rpc/sys/client/roleservice"
	"zero-fox-admin/rpc/sys/client/userservice"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	AddLog   rest.Middleware
	Redis    *redis.Redis

	//系统相关
	DeptService       deptservice.DeptService
	DictTypeService   dicttypeservice.DictTypeService
	DictItemService   dictitemservice.DictItemService
	PostService       postservice.PostService
	LoginLogService   loginlogservice.LoginLogService
	Operatelogservice operatelogservice.OperateLogService
	MenuService       menuservice.MenuService
	RoleService       roleservice.RoleService
	UserService       userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	sysClient := zrpc.MustNewClient(c.SysRpc)
	operateLogService := operatelogservice.NewOperateLogService(sysClient)
	newRedis := redis.New(c.Redis.Address, redisConfig(c))

	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,
		AddLog:   middleware.NewAddLogMiddleware(operateLogService).Handle,
		Redis:    newRedis,

		DeptService:       deptservice.NewDeptService(sysClient),
		DictTypeService:   dicttypeservice.NewDictTypeService(sysClient),
		DictItemService:   dictitemservice.NewDictItemService(sysClient),
		PostService:       postservice.NewPostService(sysClient),
		LoginLogService:   loginlogservice.NewLoginLogService(sysClient),
		Operatelogservice: operateLogService,
		MenuService:       menuservice.NewMenuService(sysClient),
		RoleService:       roleservice.NewRoleService(sysClient),
		UserService:       userservice.NewUserService(sysClient),
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
