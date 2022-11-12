package ping

import (
	"luckyProxy/common/logx"
	"luckyProxy/common/result"
	"luckyProxy/server/internal/svc"
	"net/http"
)

// Ping test function
func Ping(svcContext *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info().Msgf("hello world")
		result.HttpOk(w, "pong")
	}
}
