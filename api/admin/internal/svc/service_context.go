package svc

import (
	"zero-fox-admin/api/admin/internal/config"
	"zero-fox-admin/api/admin/internal/middleware"

	"github.com/zeromicro/go-zero/rest"

	// sys
	"zero-fox-admin/rpc/sys/client/deptservice"
	"zero-fox-admin/rpc/sys/client/dictitemservice"
	"zero-fox-admin/rpc/sys/client/dictservice"
	"zero-fox-admin/rpc/sys/client/jobservice"
	"zero-fox-admin/rpc/sys/client/loginlogservice"
	"zero-fox-admin/rpc/sys/client/menuservice"
	"zero-fox-admin/rpc/sys/client/roleservice"
	"zero-fox-admin/rpc/sys/client/syslogservice"
	"zero-fox-admin/rpc/sys/client/userservice"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware

	//系统相关
	DeptService     deptservice.DeptService
	DictService     dictservice.DictService
	DictItemService dictitemservice.DictItemService
	JobService      jobservice.JobService
	LoginLogService loginlogservice.LoginLogService
	SysLogService   syslogservice.SysLogService
	MenuService     menuservice.MenuService
	RoleService     roleservice.RoleService
	UserService     userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		CheckUrl: middleware.NewCheckUrlMiddleware().Handle,
	}
}
