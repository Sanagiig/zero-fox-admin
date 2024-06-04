package userservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sys_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleListLogic {
	return &UpdateUserRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户与角色的关联
func (l *UpdateUserRoleListLogic) UpdateUserRoleList(in *sys_client.UpdateUserRoleListReq) (*sys_client.UpdateUserRoleListResp, error) {
	// todo: add your logic here and delete this line

	return &sys_client.UpdateUserRoleListResp{}, nil
}
