package roleservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sys_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuRoleListLogic {
	return &UpdateMenuRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户与角色的关联
func (l *UpdateMenuRoleListLogic) UpdateMenuRoleList(in *sys_client.UpdateMenuRoleReq) (*sys_client.UpdateMenuRoleResp, error) {
	// todo: add your logic here and delete this line

	return &sys_client.UpdateMenuRoleResp{}, nil
}
