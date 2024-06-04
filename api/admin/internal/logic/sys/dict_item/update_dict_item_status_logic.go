package dict_item

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictItemStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictItemStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictItemStatusLogic {
	return &UpdateDictItemStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictItemStatusLogic) UpdateDictItemStatus(req *types.UpdateDictItemStatusReq) (resp *types.UpdateDictItemStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
