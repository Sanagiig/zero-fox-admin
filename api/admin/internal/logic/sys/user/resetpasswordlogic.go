package user

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReSetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReSetPasswordLogic {
	return &ReSetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReSetPasswordLogic) ReSetPassword(req *types.ReSetPasswordReq) (resp *types.ReSetPasswordResp, err error) {
	// todo: add your logic here and delete this line

	return
}
