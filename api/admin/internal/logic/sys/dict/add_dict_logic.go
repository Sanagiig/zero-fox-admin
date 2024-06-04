package dict

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictLogic {
	return &AddDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDictLogic) AddDict(req *types.AddDictReq) (resp *types.AddDictResp, err error) {
	// todo: add your logic here and delete this line

	return
}
