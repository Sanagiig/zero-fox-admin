package dept

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-fox-admin/api/admin/internal/logic/sys/dept"
	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"
)

func QueryDeptListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListDeptReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dept.NewQueryDeptListLogic(r.Context(), svcCtx)
		resp, err := l.QueryDeptList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
