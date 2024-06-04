package userservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sys_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserRoleListLogic {
	return &QueryUserRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户与角色的关联
func (l *QueryUserRoleListLogic) QueryUserRoleList(in *sys_client.QueryUserRoleListReq) (*sys_client.QueryUserRoleListResp, error) {
	// todo: add your logic here and delete this line

	return &sys_client.QueryUserRoleListResp{}, nil
}
