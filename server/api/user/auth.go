package user

import (
	"fmt"
	"github.com/rs/xid"
	"luckyProxy/common/result"
	"luckyProxy/server/cache"
	"luckyProxy/server/handler"
	"luckyProxy/server/svc"
	"net/http"
)

// Auth The current role is to generate tokens
func Auth(*svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query()
		if query != nil {
			user := handler.User{
				Name:     query.Get("username"),
				Password: query.Get("password"),
			}
			fmt.Println(user.Name + " " + user.Password)
			if user.Name != "" {
				if ok := user.Login(); !ok {
					return
				}
			}
		}
		token := xid.New().String()
		cache.ProxyInfoContainer.Add(token)
		result.HttpOk(w, token)
	}
}
