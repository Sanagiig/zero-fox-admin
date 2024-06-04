package job

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryJobListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryJobListLogic {
	return &QueryJobListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryJobListLogic) QueryJobList(req *types.ListJobReq) (resp *types.ListJobResp, err error) {
	// todo: add your logic here and delete this line

	return
}
