package proxy

import (
	"fmt"
	"io"
	pkg "luckyProxy/common"
	"net"
)

// handleUdp udp转发
func (c *Container) handleUdp(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	udpConn, err := net.ListenUDP(info.ChannelType, nil)
	if err != nil {
		return err, nil, nil
	}

	cp := &pkg.ClientProxyInfo{
		ChannelType:  info.ChannelType,
		IntranetAddr: info.Addr,
		ServerPort:   udpConn.LocalAddr().(*net.UDPAddr).Port,
	}
	info.ClientProxyInfo = cp
	info.BindListener = udpConn

	fmt.Println("udp port:", udpConn.LocalAddr().(*net.UDPAddr).Port)
	go func() {
		for {
			userConn := NewUserConn(udpConn, c, cp.Key())
			c.AddCloser(udpConn)
			c.AddUserConn(userConn)

			//指定关闭连接方式
			clean := c.CleanUserConn(userConn)

			//客户端收到
			err = userConn.OnUserConnect()
			if err != nil {
				clean()
				continue
			}

			buf := make([]byte, 1024)
			udp, addr, err := udpConn.ReadFromUDP(buf[:])
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(buf[:udp]))
			userConn.Write(buf[:udp])
			fmt.Println(addr.Port)
		}
	}()
	fmt.Println("我到这里了！")
	return nil, cp, udpConn
}
