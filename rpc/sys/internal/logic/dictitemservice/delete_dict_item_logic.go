package dictitemservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sys_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictItemLogic {
	return &DeleteDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除字典项表
func (l *DeleteDictItemLogic) DeleteDictItem(in *sys_client.DictItemDeleteReq) (*sys_client.DictItemDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &sys_client.DictItemDeleteResp{}, nil
}
