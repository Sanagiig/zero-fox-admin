package dictservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sys_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictListLogic {
	return &QueryDictListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询字典表列表
func (l *QueryDictListLogic) QueryDictList(in *sys_client.DictListReq) (*sys_client.DictListResp, error) {
	// todo: add your logic here and delete this line

	return &sys_client.DictListResp{}, nil
}
