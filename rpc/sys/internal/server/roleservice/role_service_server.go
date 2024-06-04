// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/logic/roleservice"
	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"
)

type RoleServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedRoleServiceServer
}

func NewRoleServiceServer(svcCtx *svc.ServiceContext) *RoleServiceServer {
	return &RoleServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加角色
func (s *RoleServiceServer) RoleAdd(ctx context.Context, in *sysclient.RoleAddReq) (*sysclient.RoleAddResp, error) {
	l := roleservicelogic.NewRoleAddLogic(ctx, s.svcCtx)
	return l.RoleAdd(in)
}

// 查询角色
func (s *RoleServiceServer) RoleList(ctx context.Context, in *sysclient.RoleListReq) (*sysclient.RoleListResp, error) {
	l := roleservicelogic.NewRoleListLogic(ctx, s.svcCtx)
	return l.RoleList(in)
}

// 更新角色
func (s *RoleServiceServer) RoleUpdate(ctx context.Context, in *sysclient.RoleUpdateReq) (*sysclient.RoleUpdateResp, error) {
	l := roleservicelogic.NewRoleUpdateLogic(ctx, s.svcCtx)
	return l.RoleUpdate(in)
}

// 删除角色
func (s *RoleServiceServer) RoleDelete(ctx context.Context, in *sysclient.RoleDeleteReq) (*sysclient.RoleDeleteResp, error) {
	l := roleservicelogic.NewRoleDeleteLogic(ctx, s.svcCtx)
	return l.RoleDelete(in)
}

// 查询用户与角色的关联
func (s *RoleServiceServer) QueryRoleMenuList(ctx context.Context, in *sysclient.QueryRoleMenuListReq) (*sysclient.QueryRoleMenuListResp, error) {
	l := roleservicelogic.NewQueryRoleMenuListLogic(ctx, s.svcCtx)
	return l.QueryRoleMenuList(in)
}

// 更新用户与角色的关联
func (s *RoleServiceServer) UpdateMenuRoleList(ctx context.Context, in *sysclient.UpdateMenuRoleReq) (*sysclient.UpdateMenuRoleResp, error) {
	l := roleservicelogic.NewUpdateMenuRoleListLogic(ctx, s.svcCtx)
	return l.UpdateMenuRoleList(in)
}
