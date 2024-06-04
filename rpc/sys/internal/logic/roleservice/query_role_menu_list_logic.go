package roleservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sys_client"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleMenuListLogic {
	return &QueryRoleMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户与角色的关联
func (l *QueryRoleMenuListLogic) QueryRoleMenuList(in *sys_client.QueryRoleMenuListReq) (*sys_client.QueryRoleMenuListResp, error) {
	// todo: add your logic here and delete this line

	return &sys_client.QueryRoleMenuListResp{}, nil
}
