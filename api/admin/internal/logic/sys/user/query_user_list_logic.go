package user

import (
	"context"
	"strings"

	"zero-fox-admin/api/admin/internal/common/errorx"
	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type QueryUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserListLogic {
	return &QueryUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserListLogic) QueryUserList(req *types.QueryUserListReq) (resp *types.QueryUserListResp, err error) {
	result, err := l.svcCtx.UserService.QueryUserList(l.ctx, &sysclient.QueryUserListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		DeptId:     req.DeptId,
		Email:      strings.TrimSpace(req.Email),
		Mobile:     strings.TrimSpace(req.Mobile),
		NickName:   strings.TrimSpace(req.NickName),
		UserStatus: req.UserStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询用户列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryUserListData

	for _, item := range result.List {
		list = append(list, &types.QueryUserListData{
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
		})
	}

	return &types.QueryUserListResp{
		Code:     "000000",
		Message:  "查询用户列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
