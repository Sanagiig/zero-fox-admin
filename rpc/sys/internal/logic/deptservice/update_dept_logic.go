package deptservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptLogic {
	return &UpdateDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新部门信息表
func (l *UpdateDeptLogic) UpdateDept(in *sysclient.UpdateDeptReq) (*sysclient.UpdateDeptResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdateDeptResp{}, nil
}
