package deptservicelogic

import (
	"context"
	"strconv"
	"strings"

	"zero-fox-admin/rpc/sys/internal/svc"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeptListLogic {
	return &QueryDeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询部门信息表列表
func (l *QueryDeptListLogic) QueryDeptList(in *sysclient.QueryDeptListReq) (*sysclient.QueryDeptListResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryDeptListResp{}, nil
}

// GetParentIds 获取父级ids
func GetParentIds(parentIdsStr string) []int64 {
	var parentIds []int64
	if len(parentIdsStr) > 0 {
		for _, i := range strings.Split(parentIdsStr, ",") {
			p, _ := strconv.ParseInt(i, 10, 64)
			parentIds = append(parentIds, p)
		}
	}
	return parentIds
}
