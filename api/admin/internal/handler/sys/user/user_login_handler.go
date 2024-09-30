package user

import (
	"net/http"

	"zero-fox-admin/api/admin/internal/logic/sys/user"
	"zero-fox-admin/api/admin/internal/svc"
	"zero-fox-admin/api/admin/internal/types"

	"github.com/ua-parser/uap-go/uaparser"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//获取客户端真实的ip
		//nginx的配置
		//location /api {
		//		proxy_set_header Host $host;
		//		proxy_set_header X-Real-IP $remote_addr;
		//		proxy_set_header REMOTE-HOST $remote_addr;
		//		proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		//		proxy_pass http://127.0.0.1:8888;
		//}
		userAgent := r.Header.Get("User-Agent")
		parser := uaparser.NewFromSaved()
		ua := parser.Parse(userAgent)

		browser := ua.UserAgent.Family + " " + ua.UserAgent.Major
		os := ua.Os.Family + " " + ua.Os.Major

		l := user.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req, httpx.GetRemoteAddr(r), browser, os)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
