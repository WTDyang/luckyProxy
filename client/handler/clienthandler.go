package handler

import (
	"luckyProxy/client/logic"
	"luckyProxy/client/svc"
	"luckyProxy/client/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClientHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewClientLogic(r.Context(), svcCtx)
		resp, err := l.Client(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
