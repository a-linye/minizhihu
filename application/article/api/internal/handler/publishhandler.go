package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"minizhihu/application/article/api/internal/logic"
	"minizhihu/application/article/api/internal/svc"
	"minizhihu/application/article/api/internal/types"
)

func PublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPublishLogic(r.Context(), svcCtx)
		resp, err := l.Publish(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
