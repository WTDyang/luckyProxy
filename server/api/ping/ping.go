package ping

import (
	"luckyProxy/common/logx"
	"luckyProxy/common/result"
	"luckyProxy/server/svc"
	"net/http"
)

// Ping test function
func Ping(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info().Msgf(r.Host + "ping")
		result.HttpOk(w, "pong")
	}
}
