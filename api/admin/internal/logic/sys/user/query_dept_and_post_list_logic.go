package user

import (
	"context"
	"strconv"

	"zero-fox-admin/api/admin/internal/common/errorx"
	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"
	"zero-fox-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type QueryDeptAndPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDeptAndPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDeptAndPostListLogic {
	return &QueryDeptAndPostListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryDeptAndPostList 查询所有部门和岗位
func (l *QueryDeptAndPostListLogic) QueryDeptAndPostList(req *types.QueryDeptAndPostListReq) (resp *types.QueryDeptAndPostListResp, err error) {
	//1.查询岗位列表信息
	result, err := l.svcCtx.UserService.QueryDeptAndPostList(l.ctx, &sysclient.QueryDeptAndPostListReq{})

	if err != nil {
		logc.Errorf(l.ctx, "查询所有部门和岗位失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	//2.岗位信息
	var postLists []*types.PostList
	for _, post := range result.PostListData {
		postLists = append(postLists, &types.PostList{
			Id:       post.Id,
			PostName: post.PostName,
		})
	}

	//3.部门信息
	var deptLists []*types.DeptList
	for _, dept := range result.DeptListData {
		deptLists = append(deptLists, &types.DeptList{
			Id:       dept.Id,
			DeptKey:  strconv.FormatInt(dept.Id, 10),
			Value:    strconv.FormatInt(dept.Id, 10),
			Title:    dept.DeptName,
			ParentId: dept.ParentId,
		})
	}

	return &types.QueryDeptAndPostListResp{
		Code:    "000000",
		Message: "查询成功",
		Data: types.QueryDeptAndPostListData{
			DeptList: deptLists,
			PostList: postLists,
		},
	}, nil
}
