package user

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.AddUserReq) (resp *types.AddUserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
