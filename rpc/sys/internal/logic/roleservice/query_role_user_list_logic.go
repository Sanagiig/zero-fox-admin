package roleservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryRoleUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleUserListLogic {
	return &QueryRoleUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询角色的用户关联
func (l *QueryRoleUserListLogic) QueryRoleUserList(in *sysclient.QueryRoleUserListReq) (*sysclient.QueryRoleUserListResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryRoleUserListResp{}, nil
}
