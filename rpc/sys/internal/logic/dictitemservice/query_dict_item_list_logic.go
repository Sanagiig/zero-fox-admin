package dictitemservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictItemListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictItemListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictItemListLogic {
	return &QueryDictItemListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询字典数据表列表
func (l *QueryDictItemListLogic) QueryDictItemList(in *sysclient.QueryDictItemListReq) (*sysclient.QueryDictItemListResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryDictItemListResp{}, nil
}
