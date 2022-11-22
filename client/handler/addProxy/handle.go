package addProxy

import (
	"luckyProxy/client"
	"luckyProxy/common/protocal"
)

// Handle handles addProxy
func Handle(c *client.Client, a protocal.AddProxy) {
	c.AddProxy(a)
}
