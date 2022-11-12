package handler

import (
	"luckyProxy/server/logic"
	"luckyProxy/server/svc"
	"luckyProxy/server/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewServerLogic(r.Context(), svcCtx)
		resp, err := l.Server(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
