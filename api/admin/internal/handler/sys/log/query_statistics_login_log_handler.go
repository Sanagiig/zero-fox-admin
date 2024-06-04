package log

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-fox-admin/api/admin/internal/logic/sys/log"
	"zero-fox-admin/api/admin/internal/svc"
)

func QueryStatisticsLoginLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := log.NewQueryStatisticsLoginLogLogic(r.Context(), svcCtx)
		resp, err := l.QueryStatisticsLoginLog()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
