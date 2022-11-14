package proxy

import (
	"fmt"
	"io"
	utils "luckyProxy"
	pkg "luckyProxy/common"
	"luckyProxy/common/wsx"
)

type (
	Container struct {
		// the websocket connection to client
		*wsx.Wsx
		// the Token of client
		Token string
		// closers save this client all listeners(tcp/udp/http...) associated and connections from users
		closers []io.Closer
		// UserConnMap save all user connections,
		// key is conn id
		UserConnMap map[string]*UserConn
	}
)

func NewContainer(ws *wsx.Wsx, token string) *Container {
	return &Container{Wsx: ws, Token: token, closers: []io.Closer{}, UserConnMap: make(map[string]*UserConn)}
}

// Lunch Start the local service and then generate the format of the proxy information required by the client
//
func (c Container) Lunch(infos []*pkg.ServerProxyInfo) (error, []pkg.ClientProxyInfo, []io.Closer) {
	var (
		// mapping information used to return to the client
		clientInfos []pkg.ClientProxyInfo
		// all listeners started by the current request
		listeners []io.Closer
	)

	for _, info := range infos {
		var (
			clientInfo *pkg.ClientProxyInfo
			listener   io.Closer
			err        error
		)

		switch info.ChannelType {
		case pkg.TCP:
			fmt.Println("代理tcp")
			err, clientInfo, listener = c.handleTCP(info)
		case pkg.HTTP:
			fmt.Println("代理http")
			err, clientInfo, listener = c.handlerHttp(info)
		case pkg.UDP:
			fmt.Println("代理udp")
			err, clientInfo, listener = c.handleUdp(info)
		default:
			return utils.NewError("不支持的协议类型 %s", info.ChannelType), nil, nil
		}

		if err != nil {
			return err, nil, nil
		}

		if clientInfo == nil {
			return utils.NewError("不支持的协议类型 %s", info.ChannelType), nil, nil
		}

		clientInfos = append(clientInfos, *clientInfo)
		listeners = append(listeners, listener)
	}

	return nil, clientInfos, listeners
}

// Close 关闭本地连接
func (c Container) Close() {
	c.Wsx.Close()

	for _, c := range c.closers {
		c.Close()
	}
}

// AddCloser 关闭连接的责任链
func (c *Container) AddCloser(closer io.Closer) {
	c.closers = append(c.closers, closer)
}

func (c *Container) AddUserConn(conn *UserConn) {
	c.UserConnMap[conn.Id] = conn
}

func (c *Container) GetUserConn(connId string) (*UserConn, bool) {
	userConn, ok := c.UserConnMap[connId]
	return userConn, ok
}

func (c *Container) CleanUserConn(conn *UserConn) func() {
	return func() {
		conn.conn.Close()
		delete(c.UserConnMap, conn.Id)
	}
}
