// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/logic/dictitemservice"
	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"
)

type DictItemServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedDictItemServiceServer
}

func NewDictItemServiceServer(svcCtx *svc.ServiceContext) *DictItemServiceServer {
	return &DictItemServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加字典项表
func (s *DictItemServiceServer) AddDictItem(ctx context.Context, in *sysclient.DictItemAddReq) (*sysclient.DictItemAddResp, error) {
	l := dictitemservicelogic.NewAddDictItemLogic(ctx, s.svcCtx)
	return l.AddDictItem(in)
}

// 删除字典项表
func (s *DictItemServiceServer) DeleteDictItem(ctx context.Context, in *sysclient.DictItemDeleteReq) (*sysclient.DictItemDeleteResp, error) {
	l := dictitemservicelogic.NewDeleteDictItemLogic(ctx, s.svcCtx)
	return l.DeleteDictItem(in)
}

// 更新字典项表
func (s *DictItemServiceServer) UpdateDictItem(ctx context.Context, in *sysclient.DictItemUpdateReq) (*sysclient.DictItemUpdateResp, error) {
	l := dictitemservicelogic.NewUpdateDictItemLogic(ctx, s.svcCtx)
	return l.UpdateDictItem(in)
}

// 根据条件查询单条字典项表记录
func (s *DictItemServiceServer) QueryDictItem(ctx context.Context, in *sysclient.DictItemReq) (*sysclient.DictItemResp, error) {
	l := dictitemservicelogic.NewQueryDictItemLogic(ctx, s.svcCtx)
	return l.QueryDictItem(in)
}

// 查询字典项表列表
func (s *DictItemServiceServer) QueryDictItemList(ctx context.Context, in *sysclient.DictItemListReq) (*sysclient.DictItemListResp, error) {
	l := dictitemservicelogic.NewQueryDictItemListLogic(ctx, s.svcCtx)
	return l.QueryDictItemList(in)
}
