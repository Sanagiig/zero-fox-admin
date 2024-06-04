package job

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddJobLogic {
	return &AddJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddJobLogic) AddJob(req *types.AddJobReq) (resp *types.AddJobResp, err error) {
	// todo: add your logic here and delete this line

	return
}
