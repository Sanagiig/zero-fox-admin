package dicttypeservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDictTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictTypeListLogic {
	return &QueryDictTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询字典类型表列表
func (l *QueryDictTypeListLogic) QueryDictTypeList(in *sysclient.QueryDictTypeListReq) (*sysclient.QueryDictTypeListResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryDictTypeListResp{}, nil
}
