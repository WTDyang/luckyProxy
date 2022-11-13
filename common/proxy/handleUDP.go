package proxy

import (
	"fmt"
	"io"
	pkg "luckyProxy/common"
	"net"
)

// handleUdp handler udp todo
func (c *Container) handleUdp(info *pkg.ServerProxyInfo) (error, *pkg.ClientProxyInfo, io.Closer) {
	udpConn, err := net.ListenUDP(info.ChannelType, nil)
	if err != nil {
		return err, nil, nil
	}

	fmt.Println("udp port:", udpConn.LocalAddr().(*net.UDPAddr).Port)
	go func() {
		for {
			buf := make([]byte, 1024)
			udp, addr, err := udpConn.ReadFromUDP(buf[:])
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(buf[:udp]))
			fmt.Println(addr.Port)
		}
	}()
	fmt.Println("我到这里了！")
	cp := &pkg.ClientProxyInfo{
		ChannelType:  info.ChannelType,
		IntranetAddr: info.Addr,
		ServerPort:   udpConn.LocalAddr().(*net.UDPAddr).Port,
	}
	info.ClientProxyInfo = cp
	info.BindListener = udpConn
	return nil, cp, udpConn
}
