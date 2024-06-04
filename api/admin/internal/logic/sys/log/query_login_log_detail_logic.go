package log

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLoginLogDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLoginLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLoginLogDetailLogic {
	return &QueryLoginLogDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLoginLogDetailLogic) QueryLoginLogDetail(req *types.QueryLoginLogDetailReq) (resp *types.QueryLoginLogDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
