package user

import (
	"context"

	"zero-fox-admin/api/admin/internal/common/errorx"
	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateUserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleListLogic {
	return &UpdateUserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserRoleListLogic) UpdateUserRoleList(req *types.UpdateUserRoleReq) (*types.UpdateUserRoleResp, error) {
	_, err := l.svcCtx.UserService.UpdateUserRoleList(l.ctx, &sysclient.UpdateUserRoleListReq{
		UserId:  req.UserId,
		RoleIds: req.RoleIds,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新用户与角色的关联失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateUserRoleResp{
		Code:    "000000",
		Message: "更新用户与角色的关联成功",
	}, nil
}
