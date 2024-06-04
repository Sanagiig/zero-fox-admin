package dictitemservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictItemDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictItemDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemDetailLogic {
	return &QueryDictItemDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询字典数据表详情
func (l *QueryDictItemDetailLogic) QueryDictItemDetail(in *sysclient.QueryDictItemDetailReq) (*sysclient.QueryDictItemDetailResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryDictItemDetailResp{}, nil
}
