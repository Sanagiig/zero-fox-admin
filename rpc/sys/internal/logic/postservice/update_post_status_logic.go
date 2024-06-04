package postservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostStatusLogic {
	return &UpdatePostStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新岗位管理状态
func (l *UpdatePostStatusLogic) UpdatePostStatus(in *sysclient.UpdatePostStatusReq) (*sysclient.UpdatePostStatusResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdatePostStatusResp{}, nil
}
