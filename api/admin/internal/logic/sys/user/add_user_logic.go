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

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (*types.AddUserResp, error) {
	userAddReq := sysclient.AddUserReq{
		Avatar:     req.Avatar,
		CreateBy:   l.ctx.Value("userName").(string),
		DeptId:     req.DeptId,
		Email:      req.Email,
		Mobile:     req.Mobile,
		NickName:   req.NickName,
		Remark:     req.Remark,
		UserName:   req.UserName,
		UserStatus: req.UserStatus,
		PostIds:    req.PostIds,
	}
	_, err := l.svcCtx.UserService.AddUser(l.ctx, &userAddReq)

	if err != nil {
		logc.Errorf(l.ctx, "添加用户信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddUserResp{
		Code:    "000000",
		Message: "添加用户成功",
	}, nil
}
