package log

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryStatisticsLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryStatisticsLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryStatisticsLoginLogLogic {
	return &QueryStatisticsLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryStatisticsLoginLogLogic) QueryStatisticsLoginLog() (resp *types.QueryStatisticsLoginLogResp, err error) {
	// todo: add your logic here and delete this line

	return
}
