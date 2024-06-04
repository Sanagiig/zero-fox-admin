package dict_type

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictTypeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictTypeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictTypeStatusLogic {
	return &UpdateDictTypeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateDictTypeStatusLogic) UpdateDictTypeStatus(req *types.UpdateDictTypeStatusReq) (resp *types.UpdateDictTypeStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
