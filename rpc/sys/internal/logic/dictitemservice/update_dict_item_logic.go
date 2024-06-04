package dictitemservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictItemLogic {
	return &UpdateDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新字典数据表
func (l *UpdateDictItemLogic) UpdateDictItem(in *sysclient.UpdateDictItemReq) (*sysclient.UpdateDictItemResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdateDictItemResp{}, nil
}
