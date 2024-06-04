package dicttypeservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictTypeStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictTypeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictTypeStatusLogic {
	return &UpdateDictTypeStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新字典类型表状态
func (l *UpdateDictTypeStatusLogic) UpdateDictTypeStatus(in *sysclient.UpdateDictTypeStatusReq) (*sysclient.UpdateDictTypeStatusResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdateDictTypeStatusResp{}, nil
}
