// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/logic/deptservice"
	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"
)

type DeptServiceServer struct {
	svcCtx *svc.ServiceContext
	sysclient.UnimplementedDeptServiceServer
}

func NewDeptServiceServer(svcCtx *svc.ServiceContext) *DeptServiceServer {
	return &DeptServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加部门信息表
func (s *DeptServiceServer) AddDept(ctx context.Context, in *sysclient.AddDeptReq) (*sysclient.AddDeptResp, error) {
	l := deptservicelogic.NewAddDeptLogic(ctx, s.svcCtx)
	return l.AddDept(in)
}

// 删除部门信息表
func (s *DeptServiceServer) DeleteDept(ctx context.Context, in *sysclient.DeleteDeptReq) (*sysclient.DeleteDeptResp, error) {
	l := deptservicelogic.NewDeleteDeptLogic(ctx, s.svcCtx)
	return l.DeleteDept(in)
}

// 更新部门信息表
func (s *DeptServiceServer) UpdateDept(ctx context.Context, in *sysclient.UpdateDeptReq) (*sysclient.UpdateDeptResp, error) {
	l := deptservicelogic.NewUpdateDeptLogic(ctx, s.svcCtx)
	return l.UpdateDept(in)
}

// 更新部门信息表状态
func (s *DeptServiceServer) UpdateDeptStatus(ctx context.Context, in *sysclient.UpdateDeptStatusReq) (*sysclient.UpdateDeptStatusResp, error) {
	l := deptservicelogic.NewUpdateDeptStatusLogic(ctx, s.svcCtx)
	return l.UpdateDeptStatus(in)
}

// 查询部门信息表详情
func (s *DeptServiceServer) QueryDeptDetail(ctx context.Context, in *sysclient.QueryDeptDetailReq) (*sysclient.QueryDeptDetailResp, error) {
	l := deptservicelogic.NewQueryDeptDetailLogic(ctx, s.svcCtx)
	return l.QueryDeptDetail(in)
}

// 查询部门信息表列表
func (s *DeptServiceServer) QueryDeptList(ctx context.Context, in *sysclient.QueryDeptListReq) (*sysclient.QueryDeptListResp, error) {
	l := deptservicelogic.NewQueryDeptListLogic(ctx, s.svcCtx)
	return l.QueryDeptList(in)
}
