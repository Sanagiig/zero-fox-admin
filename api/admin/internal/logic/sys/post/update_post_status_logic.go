package post

import (
	"context"

	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostStatusLogic {
	return &UpdatePostStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostStatusLogic) UpdatePostStatus(req *types.UpdatePostStatusReq) (resp *types.UpdatePostStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}