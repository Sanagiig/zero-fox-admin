// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package deptservice

import (
	"context"

	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddDeptReq                = sysclient.AddDeptReq
	AddDeptResp               = sysclient.AddDeptResp
	AddDictItemReq            = sysclient.AddDictItemReq
	AddDictItemResp           = sysclient.AddDictItemResp
	AddDictTypeReq            = sysclient.AddDictTypeReq
	AddDictTypeResp           = sysclient.AddDictTypeResp
	AddMenuReq                = sysclient.AddMenuReq
	AddMenuResp               = sysclient.AddMenuResp
	AddOperateLogReq          = sysclient.AddOperateLogReq
	AddOperateLogResp         = sysclient.AddOperateLogResp
	AddPostReq                = sysclient.AddPostReq
	AddPostResp               = sysclient.AddPostResp
	AddRoleReq                = sysclient.AddRoleReq
	AddRoleResp               = sysclient.AddRoleResp
	AddUserReq                = sysclient.AddUserReq
	AddUserResp               = sysclient.AddUserResp
	CancelAuthorizationReq    = sysclient.CancelAuthorizationReq
	CancelAuthorizationResp   = sysclient.CancelAuthorizationResp
	DeleteDeptReq             = sysclient.DeleteDeptReq
	DeleteDeptResp            = sysclient.DeleteDeptResp
	DeleteDictItemReq         = sysclient.DeleteDictItemReq
	DeleteDictItemResp        = sysclient.DeleteDictItemResp
	DeleteDictTypeReq         = sysclient.DeleteDictTypeReq
	DeleteDictTypeResp        = sysclient.DeleteDictTypeResp
	DeleteLoginLogReq         = sysclient.DeleteLoginLogReq
	DeleteLoginLogResp        = sysclient.DeleteLoginLogResp
	DeleteMenuReq             = sysclient.DeleteMenuReq
	DeleteMenuResp            = sysclient.DeleteMenuResp
	DeleteOperateLogReq       = sysclient.DeleteOperateLogReq
	DeleteOperateLogResp      = sysclient.DeleteOperateLogResp
	DeletePostReq             = sysclient.DeletePostReq
	DeletePostResp            = sysclient.DeletePostResp
	DeleteRoleReq             = sysclient.DeleteRoleReq
	DeleteRoleResp            = sysclient.DeleteRoleResp
	DeleteUserReq             = sysclient.DeleteUserReq
	DeleteUserResp            = sysclient.DeleteUserResp
	DeptListData              = sysclient.DeptListData
	DictItemListData          = sysclient.DictItemListData
	DictTypeListData          = sysclient.DictTypeListData
	InfoReq                   = sysclient.InfoReq
	InfoResp                  = sysclient.InfoResp
	LoginLogListData          = sysclient.LoginLogListData
	LoginReq                  = sysclient.LoginReq
	LoginResp                 = sysclient.LoginResp
	MenuListData              = sysclient.MenuListData
	MenuListTree              = sysclient.MenuListTree
	OperateLogListData        = sysclient.OperateLogListData
	PostListData              = sysclient.PostListData
	QueryDeptAndPostListReq   = sysclient.QueryDeptAndPostListReq
	QueryDeptAndPostListResp  = sysclient.QueryDeptAndPostListResp
	QueryDeptDetailReq        = sysclient.QueryDeptDetailReq
	QueryDeptDetailResp       = sysclient.QueryDeptDetailResp
	QueryDeptListReq          = sysclient.QueryDeptListReq
	QueryDeptListResp         = sysclient.QueryDeptListResp
	QueryDictItemDetailReq    = sysclient.QueryDictItemDetailReq
	QueryDictItemDetailResp   = sysclient.QueryDictItemDetailResp
	QueryDictItemListReq      = sysclient.QueryDictItemListReq
	QueryDictItemListResp     = sysclient.QueryDictItemListResp
	QueryDictTypeDetailReq    = sysclient.QueryDictTypeDetailReq
	QueryDictTypeDetailResp   = sysclient.QueryDictTypeDetailResp
	QueryDictTypeListReq      = sysclient.QueryDictTypeListReq
	QueryDictTypeListResp     = sysclient.QueryDictTypeListResp
	QueryLoginLogDetailReq    = sysclient.QueryLoginLogDetailReq
	QueryLoginLogDetailResp   = sysclient.QueryLoginLogDetailResp
	QueryLoginLogListReq      = sysclient.QueryLoginLogListReq
	QueryLoginLogListResp     = sysclient.QueryLoginLogListResp
	QueryMenuDetailReq        = sysclient.QueryMenuDetailReq
	QueryMenuDetailResp       = sysclient.QueryMenuDetailResp
	QueryMenuListReq          = sysclient.QueryMenuListReq
	QueryMenuListResp         = sysclient.QueryMenuListResp
	QueryOperateLogDetailReq  = sysclient.QueryOperateLogDetailReq
	QueryOperateLogDetailResp = sysclient.QueryOperateLogDetailResp
	QueryOperateLogListReq    = sysclient.QueryOperateLogListReq
	QueryOperateLogListResp   = sysclient.QueryOperateLogListResp
	QueryPostDetailReq        = sysclient.QueryPostDetailReq
	QueryPostDetailResp       = sysclient.QueryPostDetailResp
	QueryPostListReq          = sysclient.QueryPostListReq
	QueryPostListResp         = sysclient.QueryPostListResp
	QueryRoleDetailReq        = sysclient.QueryRoleDetailReq
	QueryRoleDetailResp       = sysclient.QueryRoleDetailResp
	QueryRoleListReq          = sysclient.QueryRoleListReq
	QueryRoleListResp         = sysclient.QueryRoleListResp
	QueryRoleMenuListReq      = sysclient.QueryRoleMenuListReq
	QueryRoleMenuListResp     = sysclient.QueryRoleMenuListResp
	QueryRoleUserListReq      = sysclient.QueryRoleUserListReq
	QueryRoleUserListResp     = sysclient.QueryRoleUserListResp
	QueryUserDetailReq        = sysclient.QueryUserDetailReq
	QueryUserDetailResp       = sysclient.QueryUserDetailResp
	QueryUserListReq          = sysclient.QueryUserListReq
	QueryUserListResp         = sysclient.QueryUserListResp
	QueryUserRoleListReq      = sysclient.QueryUserRoleListReq
	QueryUserRoleListResp     = sysclient.QueryUserRoleListResp
	ReSetPasswordReq          = sysclient.ReSetPasswordReq
	ReSetPasswordResp         = sysclient.ReSetPasswordResp
	RoleListData              = sysclient.RoleListData
	UpdateDeptReq             = sysclient.UpdateDeptReq
	UpdateDeptResp            = sysclient.UpdateDeptResp
	UpdateDeptStatusReq       = sysclient.UpdateDeptStatusReq
	UpdateDeptStatusResp      = sysclient.UpdateDeptStatusResp
	UpdateDictItemReq         = sysclient.UpdateDictItemReq
	UpdateDictItemResp        = sysclient.UpdateDictItemResp
	UpdateDictItemStatusReq   = sysclient.UpdateDictItemStatusReq
	UpdateDictItemStatusResp  = sysclient.UpdateDictItemStatusResp
	UpdateDictTypeReq         = sysclient.UpdateDictTypeReq
	UpdateDictTypeResp        = sysclient.UpdateDictTypeResp
	UpdateDictTypeStatusReq   = sysclient.UpdateDictTypeStatusReq
	UpdateDictTypeStatusResp  = sysclient.UpdateDictTypeStatusResp
	UpdateMenuReq             = sysclient.UpdateMenuReq
	UpdateMenuResp            = sysclient.UpdateMenuResp
	UpdateMenuRoleReq         = sysclient.UpdateMenuRoleReq
	UpdateMenuRoleResp        = sysclient.UpdateMenuRoleResp
	UpdateMenuStatusReq       = sysclient.UpdateMenuStatusReq
	UpdateMenuStatusResp      = sysclient.UpdateMenuStatusResp
	UpdatePostReq             = sysclient.UpdatePostReq
	UpdatePostResp            = sysclient.UpdatePostResp
	UpdatePostStatusReq       = sysclient.UpdatePostStatusReq
	UpdatePostStatusResp      = sysclient.UpdatePostStatusResp
	UpdateRoleReq             = sysclient.UpdateRoleReq
	UpdateRoleResp            = sysclient.UpdateRoleResp
	UpdateRoleStatusReq       = sysclient.UpdateRoleStatusReq
	UpdateRoleStatusResp      = sysclient.UpdateRoleStatusResp
	UpdateUserReq             = sysclient.UpdateUserReq
	UpdateUserResp            = sysclient.UpdateUserResp
	UpdateUserRoleListReq     = sysclient.UpdateUserRoleListReq
	UpdateUserRoleListResp    = sysclient.UpdateUserRoleListResp
	UpdateUserStatusReq       = sysclient.UpdateUserStatusReq
	UpdateUserStatusResp      = sysclient.UpdateUserStatusResp
	UserListData              = sysclient.UserListData

	DeptService interface {
		// 添加部门信息表
		AddDept(ctx context.Context, in *AddDeptReq, opts ...grpc.CallOption) (*AddDeptResp, error)
		// 删除部门信息表
		DeleteDept(ctx context.Context, in *DeleteDeptReq, opts ...grpc.CallOption) (*DeleteDeptResp, error)
		// 更新部门信息表
		UpdateDept(ctx context.Context, in *UpdateDeptReq, opts ...grpc.CallOption) (*UpdateDeptResp, error)
		// 更新部门信息表状态
		UpdateDeptStatus(ctx context.Context, in *UpdateDeptStatusReq, opts ...grpc.CallOption) (*UpdateDeptStatusResp, error)
		// 查询部门信息表详情
		QueryDeptDetail(ctx context.Context, in *QueryDeptDetailReq, opts ...grpc.CallOption) (*QueryDeptDetailResp, error)
		// 查询部门信息表列表
		QueryDeptList(ctx context.Context, in *QueryDeptListReq, opts ...grpc.CallOption) (*QueryDeptListResp, error)
	}

	defaultDeptService struct {
		cli zrpc.Client
	}
)

func NewDeptService(cli zrpc.Client) DeptService {
	return &defaultDeptService{
		cli: cli,
	}
}

// 添加部门信息表
func (m *defaultDeptService) AddDept(ctx context.Context, in *AddDeptReq, opts ...grpc.CallOption) (*AddDeptResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.AddDept(ctx, in, opts...)
}

// 删除部门信息表
func (m *defaultDeptService) DeleteDept(ctx context.Context, in *DeleteDeptReq, opts ...grpc.CallOption) (*DeleteDeptResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.DeleteDept(ctx, in, opts...)
}

// 更新部门信息表
func (m *defaultDeptService) UpdateDept(ctx context.Context, in *UpdateDeptReq, opts ...grpc.CallOption) (*UpdateDeptResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.UpdateDept(ctx, in, opts...)
}

// 更新部门信息表状态
func (m *defaultDeptService) UpdateDeptStatus(ctx context.Context, in *UpdateDeptStatusReq, opts ...grpc.CallOption) (*UpdateDeptStatusResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.UpdateDeptStatus(ctx, in, opts...)
}

// 查询部门信息表详情
func (m *defaultDeptService) QueryDeptDetail(ctx context.Context, in *QueryDeptDetailReq, opts ...grpc.CallOption) (*QueryDeptDetailResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.QueryDeptDetail(ctx, in, opts...)
}

// 查询部门信息表列表
func (m *defaultDeptService) QueryDeptList(ctx context.Context, in *QueryDeptListReq, opts ...grpc.CallOption) (*QueryDeptListResp, error) {
	client := sysclient.NewDeptServiceClient(m.cli.Conn())
	return client.QueryDeptList(ctx, in, opts...)
}