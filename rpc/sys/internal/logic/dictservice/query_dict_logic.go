package dictservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sys_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictLogic {
	return &QueryDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据条件查询单条字典表记录
func (l *QueryDictLogic) QueryDict(in *sys_client.DictReq) (*sys_client.DictResp, error) {
	// todo: add your logic here and delete this line

	return &sys_client.DictResp{}, nil
}
