package menuservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuListLogic {
	return &QueryMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询菜单信息表列表
func (l *QueryMenuListLogic) QueryMenuList(in *sysclient.QueryMenuListReq) (*sysclient.QueryMenuListResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryMenuListResp{}, nil
}
