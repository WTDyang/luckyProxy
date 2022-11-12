package user

import (
	"github.com/rs/xid"
	"luckyProxy/common/result"
	"luckyProxy/server/cache"
	"luckyProxy/server/svc"
	"net/http"
)

// Auth The current role is to generate tokens
func Auth(*svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := xid.New().String()

		cache.ProxyInfoContainer.Add(token)
		result.HttpOk(w, token)
	}
}
