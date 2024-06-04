package dict

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictListLogic {
	return &QueryDictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictListLogic) QueryDictList(req *types.ListDictReq) (resp *types.ListDictResp, err error) {
	// todo: add your logic here and delete this line

	return
}
