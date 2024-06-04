package deptservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeptStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeptStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeptStatusLogic {
	return &UpdateDeptStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新部门信息表状态
func (l *UpdateDeptStatusLogic) UpdateDeptStatus(in *sysclient.UpdateDeptStatusReq) (*sysclient.UpdateDeptStatusResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.UpdateDeptStatusResp{}, nil
}
