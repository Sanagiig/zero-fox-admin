package dict_item

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictItemDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemDetailLogic {
	return &QueryDictItemDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictItemDetailLogic) QueryDictItemDetail(req *types.QueryDictItemDetailReq) (resp *types.QueryDictItemDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
