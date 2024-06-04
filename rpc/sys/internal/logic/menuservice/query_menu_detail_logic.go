package menuservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuDetailLogic {
	return &QueryMenuDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询菜单信息表详情
func (l *QueryMenuDetailLogic) QueryMenuDetail(in *sysclient.QueryMenuDetailReq) (*sysclient.QueryMenuDetailResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryMenuDetailResp{}, nil
}
