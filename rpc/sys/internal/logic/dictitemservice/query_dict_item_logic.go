package dictitemservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sys_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemLogic {
	return &QueryDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据条件查询单条字典项表记录
func (l *QueryDictItemLogic) QueryDictItem(in *sys_client.DictItemReq) (*sys_client.DictItemResp, error) {
	// todo: add your logic here and delete this line

	return &sys_client.DictItemResp{}, nil
}
