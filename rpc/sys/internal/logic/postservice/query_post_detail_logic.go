package postservicelogic

import (
	"context"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPostDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPostDetailLogic {
	return &QueryPostDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询岗位管理详情
func (l *QueryPostDetailLogic) QueryPostDetail(in *sysclient.QueryPostDetailReq) (*sysclient.QueryPostDetailResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryPostDetailResp{}, nil
}
