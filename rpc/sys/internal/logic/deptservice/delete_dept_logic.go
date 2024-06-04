package deptservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeptLogic {
	return &DeleteDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除部门信息表
func (l *DeleteDeptLogic) DeleteDept(in *sysclient.DeleteDeptReq) (*sysclient.DeleteDeptResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.DeleteDeptResp{}, nil
}
