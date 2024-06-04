package dictservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sys_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictLogic {
	return &AddDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加字典表
func (l *AddDictLogic) AddDict(in *sys_client.DictAddReq) (*sys_client.DictAddResp, error) {
	// todo: add your logic here and delete this line

	return &sys_client.DictAddResp{}, nil
}
