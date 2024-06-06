// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/logic/menuservice"
	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"
)

type MenuServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedMenuServiceServer
}

func NewMenuServiceServer(svcCtx *svc.ServiceContext) *MenuServiceServer {
	return &MenuServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加菜单信息表
func (s *MenuServiceServer) AddMenu(ctx context.Context, in *sysclient.AddMenuReq) (*sysclient.AddMenuResp, error) {
	l := menuservicelogic.NewAddMenuLogic(ctx, s.svcCtx)
	return l.AddMenu(in)
}

// 删除菜单信息表
func (s *MenuServiceServer) DeleteMenu(ctx context.Context, in *sysclient.DeleteMenuReq) (*sysclient.DeleteMenuResp, error) {
	l := menuservicelogic.NewDeleteMenuLogic(ctx, s.svcCtx)
	return l.DeleteMenu(in)
}

// 更新菜单信息表
func (s *MenuServiceServer) UpdateMenu(ctx context.Context, in *sysclient.UpdateMenuReq) (*sysclient.UpdateMenuResp, error) {
	l := menuservicelogic.NewUpdateMenuLogic(ctx, s.svcCtx)
	return l.UpdateMenu(in)
}

// 更新菜单信息表状态
func (s *MenuServiceServer) UpdateMenuStatus(ctx context.Context, in *sysclient.UpdateMenuStatusReq) (*sysclient.UpdateMenuStatusResp, error) {
	l := menuservicelogic.NewUpdateMenuStatusLogic(ctx, s.svcCtx)
	return l.UpdateMenuStatus(in)
}

// 查询菜单信息表详情
func (s *MenuServiceServer) QueryMenuDetail(ctx context.Context, in *sysclient.QueryMenuDetailReq) (*sysclient.QueryMenuDetailResp, error) {
	l := menuservicelogic.NewQueryMenuDetailLogic(ctx, s.svcCtx)
	return l.QueryMenuDetail(in)
}

// 查询菜单信息表列表
func (s *MenuServiceServer) QueryMenuList(ctx context.Context, in *sysclient.QueryMenuListReq) (*sysclient.QueryMenuListResp, error) {
	l := menuservicelogic.NewQueryMenuListLogic(ctx, s.svcCtx)
	return l.QueryMenuList(in)
}