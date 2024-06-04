package dicttypeservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDictTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDictTypeLogic {
	return &DeleteDictTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除字典类型表
func (l *DeleteDictTypeLogic) DeleteDictType(in *sysclient.DeleteDictTypeReq) (*sysclient.DeleteDictTypeResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.DeleteDictTypeResp{}, nil
}
