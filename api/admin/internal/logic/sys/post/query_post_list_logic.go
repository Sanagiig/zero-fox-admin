package post

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostListLogic {
	return &QueryPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryPostListLogic) QueryPostList(req *types.QueryPostListReq) (resp *types.QueryPostListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
