package deptservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDeptLogic {
	return &AddDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加部门信息表
func (l *AddDeptLogic) AddDept(in *sysclient.AddDeptReq) (*sysclient.AddDeptResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.AddDeptResp{}, nil
}
