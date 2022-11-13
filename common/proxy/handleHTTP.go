package proxy

import (
	"io"
	pkg "luckyProxy/common"
)

func (c *Container) handlerHttp(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	// todo handler http
	return c.handleTCP(info)
}
