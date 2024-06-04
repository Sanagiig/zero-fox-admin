package dicttypeservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictTypeLogic {
	return &UpdateDictTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新字典类型表
func (l *UpdateDictTypeLogic) UpdateDictType(in *sysclient.UpdateDictTypeReq) (*sysclient.UpdateDictTypeResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdateDictTypeResp{}, nil
}
