package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"minizhihu/application/article/api/internal/logic"
	"minizhihu/application/article/api/internal/svc"
)

func UploadCoverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadCoverLogic(r.Context(), svcCtx)
		resp, err := l.UploadCover()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
