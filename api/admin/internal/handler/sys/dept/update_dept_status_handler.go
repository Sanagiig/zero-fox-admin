package dept

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-fox-admin/api/admin/internal/logic/sys/dept"
	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"
)

func UpdateDeptStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateDeptStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dept.NewUpdateDeptStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateDeptStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
