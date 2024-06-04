package dicttypeservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDictTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDictTypeLogic {
	return &AddDictTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加字典类型表
func (l *AddDictTypeLogic) AddDictType(in *sysclient.AddDictTypeReq) (*sysclient.AddDictTypeResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.AddDictTypeResp{}, nil
}
