package proxy

import (
	"errors"
	"fmt"
	"io"
	pkg "luckyProxy/common"
	"luckyProxy/common/logx"
	"net"
	"strings"
)

// handleUdp udp转发
func (c *Container) handleUdp(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	//监听本地udp端口
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
		fmt.Printf("开始监听%s\n", udpConn.LocalAddr())
		for {
			userConn := NewUserConn(udpConn, c, cp.Key())
			c.AddCloser(udpConn)
			c.AddUserConn(userConn)

			//指定关闭连接方式
			clean := c.CleanUserConn(userConn)

			//客户端与连接绑定
			err = userConn.OnUserConnect()
			if err != nil {
				fmt.Printf("都是我的错%v", err)
				clean()
				continue
			}
			//接收信息
			buf := make([]byte, 1024)
			n, _, err := udpConn.ReadFromUDP(buf[:])
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			fmt.Println(string(buf[:n]))
			//发送信息给客户端
			err = udpSend(info.Addr, buf[:n])
			if err != nil {
				logx.Err(err)
				continue
			}
			fmt.Printf("发送结束->%s\n", info.Addr)
		}
	}()
	fmt.Println("我到这里了！")
	return nil, cp, udpConn
}
func udpSend(Addr string, buf []byte) error {
	index := strings.IndexAny(Addr, ":")
	if index == -1 {
		return errors.New("格式错误")
	}
	fmt.Println(Addr[:index])
	if Addr[:index] == "localhost" {
		Addr = "127.0.0.1" + Addr[index:]
	}
	fmt.Println(Addr, string(buf))
	socket, err := net.Dial("udp", Addr)
	if err != nil {
		return err
	}
	defer func(socket net.Conn) {
		err := socket.Close()
		if err != nil {
			logx.Err(err)
		}
	}(socket)
	_, err = socket.Write(buf)
	return err
}
