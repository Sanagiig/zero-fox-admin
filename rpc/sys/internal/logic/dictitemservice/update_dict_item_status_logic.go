package dictitemservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDictItemStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDictItemStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictItemStatusLogic {
	return &UpdateDictItemStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新字典数据表状态
func (l *UpdateDictItemStatusLogic) UpdateDictItemStatus(in *sysclient.UpdateDictItemStatusReq) (*sysclient.UpdateDictItemStatusResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdateDictItemStatusResp{}, nil
}
