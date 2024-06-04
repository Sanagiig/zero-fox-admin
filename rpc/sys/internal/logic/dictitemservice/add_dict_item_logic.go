package dictitemservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictItemLogic {
	return &AddDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加字典数据表
func (l *AddDictItemLogic) AddDictItem(in *sysclient.AddDictItemReq) (*sysclient.AddDictItemResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.AddDictItemResp{}, nil
}
