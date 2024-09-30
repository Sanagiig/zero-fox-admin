package user

import (
	"context"

	"zero-fox-admin/api/admin/internal/common/errorx"
	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"
	"zero-fox-admin/rpc/sys/sysclient"

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

func (l *ReSetPasswordLogic) ReSetPassword(req *types.ReSetPasswordReq) (*types.ReSetPasswordResp, error) {
	_, err := l.svcCtx.UserService.ReSetPassword(l.ctx, &sysclient.ReSetPasswordReq{
		Id:       req.UserId,
		UpdateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		return nil, errorx.NewDefaultError("重置用户密码异常")
	}

	return &types.ReSetPasswordResp{
		Code:    "000000",
		Message: "重置用户密码成功",
	}, nil
}
