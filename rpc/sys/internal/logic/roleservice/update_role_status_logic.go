package roleservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleStatusLogic {
	return &UpdateRoleStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色信息表状态
func (l *UpdateRoleStatusLogic) UpdateRoleStatus(in *sysclient.UpdateRoleStatusReq) (*sysclient.UpdateRoleStatusResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdateRoleStatusResp{}, nil
}
