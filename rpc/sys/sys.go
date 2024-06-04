package main

import (
	"flag"
	"fmt"

	"zero-fox-admin/rpc/sys/internal/config"
	deptserviceServer "zero-fox-admin/rpc/sys/internal/server/deptservice"
	dictitemserviceServer "zero-fox-admin/rpc/sys/internal/server/dictitemservice"
	dictserviceServer "zero-fox-admin/rpc/sys/internal/server/dictservice"
	jobserviceServer "zero-fox-admin/rpc/sys/internal/server/jobservice"
	loginlogserviceServer "zero-fox-admin/rpc/sys/internal/server/loginlogservice"
	menuserviceServer "zero-fox-admin/rpc/sys/internal/server/menuservice"
	roleserviceServer "zero-fox-admin/rpc/sys/internal/server/roleservice"
	syslogserviceServer "zero-fox-admin/rpc/sys/internal/server/syslogservice"
	userserviceServer "zero-fox-admin/rpc/sys/internal/server/userservice"
	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sysclient.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		sysclient.RegisterRoleServiceServer(grpcServer, roleserviceServer.NewRoleServiceServer(ctx))
		sysclient.RegisterMenuServiceServer(grpcServer, menuserviceServer.NewMenuServiceServer(ctx))
		sysclient.RegisterDictServiceServer(grpcServer, dictserviceServer.NewDictServiceServer(ctx))
		sysclient.RegisterDictItemServiceServer(grpcServer, dictitemserviceServer.NewDictItemServiceServer(ctx))
		sysclient.RegisterDeptServiceServer(grpcServer, deptserviceServer.NewDeptServiceServer(ctx))
		sysclient.RegisterLoginLogServiceServer(grpcServer, loginlogserviceServer.NewLoginLogServiceServer(ctx))
		sysclient.RegisterSysLogServiceServer(grpcServer, syslogserviceServer.NewSysLogServiceServer(ctx))
		sysclient.RegisterJobServiceServer(grpcServer, jobserviceServer.NewJobServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
