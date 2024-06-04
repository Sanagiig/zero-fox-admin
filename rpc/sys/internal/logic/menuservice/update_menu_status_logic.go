package menuservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuStatusLogic {
	return &UpdateMenuStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新菜单信息表状态
func (l *UpdateMenuStatusLogic) UpdateMenuStatus(in *sysclient.UpdateMenuStatusReq) (*sysclient.UpdateMenuStatusResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdateMenuStatusResp{}, nil
}
