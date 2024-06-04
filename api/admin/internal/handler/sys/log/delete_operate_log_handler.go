package log

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-fox-admin/api/admin/internal/logic/sys/log"
	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"
)

func DeleteOperateLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteOperateLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := log.NewDeleteOperateLogLogic(r.Context(), svcCtx)
		resp, err := l.DeleteOperateLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
