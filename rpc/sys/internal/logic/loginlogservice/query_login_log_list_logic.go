package loginlogservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLoginLogListLogic {
	return &QueryLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询系统登录日志表列表
func (l *QueryLoginLogListLogic) QueryLoginLogList(in *sysclient.QueryLoginLogListReq) (*sysclient.QueryLoginLogListResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryLoginLogListResp{}, nil
}
