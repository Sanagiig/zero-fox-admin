package loginlogservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLoginLogLogic {
	return &DeleteLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除系统登录日志表
func (l *DeleteLoginLogLogic) DeleteLoginLog(in *sysclient.DeleteLoginLogReq) (*sysclient.DeleteLoginLogResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.DeleteLoginLogResp{}, nil
}
