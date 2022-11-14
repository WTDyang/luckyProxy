package proxy

import (
	"io"
	pkg "luckyProxy/common"
	"luckyProxy/common/logx"
	"net"
	"strings"
)

func (c *Container) handleTCP(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	tcp, err := net.ListenTCP(info.ChannelType, nil)
	if err != nil {
		return err, nil, nil
	}

	c.AddCloser(tcp)
	serverPort := tcp.Addr().(*net.TCPAddr).Port

	cp := &pkg.ClientProxyInfo{
		ChannelType:  info.ChannelType,
		IntranetAddr: info.Addr,
		ServerPort:   serverPort,
	}
	info.ClientProxyInfo = cp
	info.BindListener = tcp

	//开启线程
	go func() {
		for {
			// 首先和客户端建立tcp连接
			conn, err := tcp.AcceptTCP()
			if err != nil {
				if strings.ContainsAny("use of closed network connection", err.Error()) || strings.ContainsAny("EOF", err.Error()) {
					return
				}

				logx.Err(err).Str("channelType", info.ChannelType).Msg("accept user connection")
				return
			}

			//对建立的tcp连接进行包装
			userConn := NewUserConn(conn, c, cp.Key())
			c.AddCloser(conn)
			c.AddUserConn(userConn)

			//指定关闭连接方式
			clean := c.CleanUserConn(userConn)

			//客户端收到
			err = userConn.OnUserConnect()
			if err != nil {
				clean()
				continue
			}
			
			info.BindUserConn = append(info.BindUserConn, userConn.conn)
			go userConn.StartRead(clean)
			go userConn.StartWrite(clean)
		}
	}()

	return nil, cp, tcp
}
