package svc

import (
	"github.com/gorilla/websocket"
	"luckyProxy/server/internal/config"
	"net/http"
)

type (
	ServiceContext struct {
		Config     config.Config
		WsUpgrader websocket.Upgrader
	}
)

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		WsUpgrader: newWsUpgrader(c),
	}
}

// newWsUpgrader new websocket.Upgrader instance
func newWsUpgrader(c config.Config) websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}
