package user

import (
	"context"
	"strconv"
	"strings"

	"zero-fox-admin/api/admin/internal/common/errorx"
	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq, ip, browser, os string) (*types.LoginResp, error) {
	resp, err := l.svcCtx.UserService.Login(l.ctx, &sysclient.LoginReq{
		Account:   strings.TrimSpace(req.Account),
		Password:  strings.TrimSpace(req.Password),
		IpAddress: ip,
		Browser:   browser,
		Os:        os,
	})

	if err != nil {
		logc.Errorf(l.ctx, "用户登录：%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	//把能访问的url存在在redis，在middleware中校验
	key := "zero:mall:token:" + strconv.FormatInt(resp.Id, 10)
	err = l.svcCtx.Redis.Set(key, strings.Join(resp.ApiUrls, ","))
	if err != nil {
		logc.Errorf(l.ctx, "设置用户：%s,权限到redis异常: %+v", resp.UserName, err)
	}
	return &types.LoginResp{
		Code:    "000000",
		Message: "登录成功",
		Data: types.LoginData{
			AccessToken: resp.AccessToken,
		},
	}, nil
}
