package role

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryRoleUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleUserListLogic {
	return &QueryRoleUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryRoleUserListLogic) QueryRoleUserList(req *types.QueryRoleUserListReq) (resp *types.QueryRoleUserListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
