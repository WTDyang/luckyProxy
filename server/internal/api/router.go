package api

import (
	"github.com/zeromicro/go-zero/rest"
	"luckyProxy/server/internal/api/ping"
	"luckyProxy/server/internal/api/proxy"
	"luckyProxy/server/internal/api/user"
	"luckyProxy/server/internal/api/ws"
	"luckyProxy/server/internal/svc"
	"net/http"
)

func MountRouters(s *rest.Server, svcContext *svc.ServiceContext) {
	s.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/accept",
				Handler: ws.Accept(svcContext),
			},
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: ping.Ping(svcContext),
			},
		},
	)

	s.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/auth",
			Handler: user.Auth(svcContext),
		},
	},
		rest.WithPrefix("/user"),
	)

	s.AddRoutes([]rest.Route{
		{
			Method:  http.MethodPost,
			Path:    "/add/:token",
			Handler: proxy.AddProxy(svcContext),
		},
		{
			Method:  http.MethodPost,
			Path:    "/remove/:token",
			Handler: proxy.RemoveProxy(svcContext),
		},
	},
		rest.WithPrefix("/proxy"),
	)
}
