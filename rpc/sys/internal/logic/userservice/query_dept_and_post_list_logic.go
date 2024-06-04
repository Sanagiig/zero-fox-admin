package userservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDeptAndPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDeptAndPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeptAndPostListLogic {
	return &QueryDeptAndPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询所有部门和岗位
func (l *QueryDeptAndPostListLogic) QueryDeptAndPostList(in *sysclient.QueryDeptAndPostListReq) (*sysclient.QueryDeptAndPostListResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryDeptAndPostListResp{}, nil
}
