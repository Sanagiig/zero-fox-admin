package operatelogservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOperateLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOperateLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOperateLogLogic {
	return &DeleteOperateLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除系统操作日志表
func (l *DeleteOperateLogLogic) DeleteOperateLog(in *sysclient.DeleteOperateLogReq) (*sysclient.DeleteOperateLogResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.DeleteOperateLogResp{}, nil
}
