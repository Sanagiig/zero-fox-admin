package deptservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDeptDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDeptDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeptDetailLogic {
	return &QueryDeptDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询部门信息表详情
func (l *QueryDeptDetailLogic) QueryDeptDetail(in *sysclient.QueryDeptDetailReq) (*sysclient.QueryDeptDetailResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryDeptDetailResp{}, nil
}
