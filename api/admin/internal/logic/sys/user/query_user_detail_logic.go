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

type QueryUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserDetailLogic {
	return &QueryUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserDetailLogic) QueryUserDetail(req *types.QueryUserDetailReq) (resp *types.QueryUserDetailResp, err error) {
	item, err := l.svcCtx.UserService.QueryUserDetail(l.ctx, &sysclient.QueryUserDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户详情,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryUserDetailData{
		Avatar:     item.Avatar,
		CreateBy:   item.CreateBy,
		CreateTime: item.CreateTime,
		DeptId:     item.DeptId,
		Email:      item.Email,
		Id:         item.Id,
		LoginIp:    item.LoginIp,
		LoginTime:  item.LoginTime,
		Mobile:     item.Mobile,
		NickName:   item.NickName,
		Remark:     item.Remark,
		UpdateBy:   item.UpdateBy,
		UpdateTime: item.UpdateTime,
		UserName:   item.UserName,
		UserStatus: item.UserStatus,
		PostIds:    item.PostIds,
	}

	return &types.QueryUserDetailResp{
		Code:    "000000",
		Message: "查询用户详情成功",
		Data:    data,
	}, nil
}
