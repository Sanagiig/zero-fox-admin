package user

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
