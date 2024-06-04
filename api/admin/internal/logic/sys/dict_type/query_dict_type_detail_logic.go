package dict_type

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictTypeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictTypeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictTypeDetailLogic {
	return &QueryDictTypeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictTypeDetailLogic) QueryDictTypeDetail(req *types.QueryDictTypeDetailReq) (resp *types.QueryDictTypeDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
