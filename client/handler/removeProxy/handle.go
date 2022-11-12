package removeProxy

import (
	"luckyProxy/client"
	"luckyProxy/common/protocal"
)

func Handle(c *client.Client, removeProxy protocal.RemoveProxy) {
	c.RemoveProxy(removeProxy)

	c.CloseInternet(removeProxy.Proxy)
}
