package log

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOperateLogDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOperateLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOperateLogDetailLogic {
	return &QueryOperateLogDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOperateLogDetailLogic) QueryOperateLogDetail(req *types.QueryOperateLogDetailReq) (resp *types.QueryOperateLogDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
