package dict_type

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDictTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictTypeLogic {
	return &AddDictTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDictTypeLogic) AddDictType(req *types.AddDictTypeReq) (resp *types.AddDictTypeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
