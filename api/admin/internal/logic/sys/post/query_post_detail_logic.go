package post

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPostDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostDetailLogic {
	return &QueryPostDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPostDetailLogic) QueryPostDetail(req *types.QueryPostDetailReq) (resp *types.QueryPostDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
