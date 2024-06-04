package roleservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelAuthorizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelAuthorizationLogic {
	return &CancelAuthorizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消授权/确认授权
func (l *CancelAuthorizationLogic) CancelAuthorization(in *sysclient.CancelAuthorizationReq) (*sysclient.CancelAuthorizationResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.CancelAuthorizationResp{}, nil
}
