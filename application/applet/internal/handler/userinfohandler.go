package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"minizhihu/application/applet/internal/logic"
	"minizhihu/application/applet/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
