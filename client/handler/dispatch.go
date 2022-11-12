package handler

import (
	"luckyProxy/client"
	"luckyProxy/client/internal/handler/addProxy"
	"luckyProxy/client/internal/handler/removeProxy"
	"luckyProxy/client/internal/handler/userConnect"
	"luckyProxy/client/internal/handler/userRequest"
	"luckyProxy/common/logx"
	"luckyProxy/common/protocal"
)

func Dispatch(c *client.Client) func(bytes []byte) {
	return func(bytes []byte) {
		burst, err := protocal.Decode(bytes)
		if err != nil {
			logx.Err(err).Msg("decode burst")
			return
		}

		switch burst.Type {
		case protocal.AddProxyType:
			addProxy.Handle(c, burst.AddProxy)
		case protocal.RemoveProxyType:
			removeProxy.Handle(c, burst.RemoveProxy)
		case protocal.UserConnectType:
			userConnect.Handle(c, burst.UserConnect)
		case protocal.UserRequestType:
			userRequest.Handle(c, burst.UserRequest)
		}
	}
}
