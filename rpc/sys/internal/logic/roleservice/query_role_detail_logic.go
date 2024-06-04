package roleservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleDetailLogic {
	return &QueryRoleDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询角色信息表详情
func (l *QueryRoleDetailLogic) QueryRoleDetail(in *sysclient.QueryRoleDetailReq) (*sysclient.QueryRoleDetailResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryRoleDetailResp{}, nil
}
