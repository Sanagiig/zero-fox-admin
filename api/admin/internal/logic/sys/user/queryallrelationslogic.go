package user

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAllRelationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAllRelationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAllRelationsLogic {
	return &QueryAllRelationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAllRelationsLogic) QueryAllRelations(req *types.QueryAllRelationsReq) (resp *types.QueryAllRelationsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
