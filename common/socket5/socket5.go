package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type Socket5 struct {
	Type string
	ip   string
	port string
}

const (
	socks5Ver = 0x05

	noAuthenticationRequired = 0x00
	GSSAPI                   = 0x01
	nameAndPwd               = 0x02

	cmdBind      = 0x01
	bind         = 0x02
	UDPAssociate = 0x03

	atypIPV4  = 0x01
	atypeHOST = 0x03
	atypeIPV6 = 0x04
)

const (
	succeeded = iota
	generalSOCKSServerFailure
	connectionNotAllowedByRuleset
	NetworkUnreachable
	HostUnreachable
	ConnectionRefused
	TTLExpired
	CommandNotSupported
	AddressTypeNotSupported
	unassigned
)

func main() {
	Run("tcp", "127.0.0.1", "1080")
}

func Run(Type string, ip string, port string) error {
	s := &Socket5{
		Type: Type,
		ip:   ip,
		port: port,
	}
	server, err := net.Listen("tcp", s.GetAddr())
	if err != nil {
		return fmt.Errorf("socket5 error[%e] in init time ", err)
	}
	for {
		//拿到连接
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		//处理连接
		go func() {
			err := s.Handler(client)
			if err != nil {
				log.Printf("client handler error : %v", err)
			}
		}()

	}
}

func (s *Socket5) GetAddr() string {
	return s.ip + ":" + s.port
}

func (s *Socket5) Handler(conn net.Conn) error {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//进行认证 协商
	err := s.Auth(reader, conn)
	if err != nil {
		return fmt.Errorf("client %v auth failed:%v", conn.RemoteAddr(), err)
	}
	//建立连接 开始转发
	err = s.Connect(reader, conn)
	if err != nil {
		return fmt.Errorf("client %v auth failed:%v", conn.RemoteAddr(), err)
	}
	return nil
}

//Auth 认证过程
func (s *Socket5) Auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	//读取一个字节，也就是VER
	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	//读取下一个字节，也就是NMETHODS
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	//读取支持方法列表
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}

	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	//返回自己支持的方法 (不需要验证)
	_, err = conn.Write([]byte{socks5Ver, noAuthenticationRequired})
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil
}

func (s *Socket5) Connect(reader *bufio.Reader, conn net.Conn) (err error) {
	//客户端报文格式
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 版本号，socks5的值为0x05
	// CMD CONNECT请求CONNECT 0x01 连接 BIND 0x02 端口监听(也就是在Server上监听一个端口) UDP ASSOCIATE 0x03 使用UDP
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名
	// DST.ADDR 一个可变长度的值
	// DST.PORT 目标端口，固定2个字节

	//读取固定四位
	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header failed:%w", err)
	}
	ver, cmd, atyp := buf[0], buf[1], buf[3]
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supported cmd:%v", ver)
	}
	//读取访问地址
	addr := ""
	Atype := atypIPV4
	ip := make([]byte, 4)
	switch atyp {
	case atypIPV4:
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
		ip[0], ip[1], ip[2], ip[3] = buf[0], buf[1], buf[2], buf[3]
	case atypeIPV6:
		return errors.New("IPv6: no supported yet")
		Atype = atypeIPV6
	case atypeHOST:
		//域名类型第一个字节表示域名的长度
		Atype = atypeHOST
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read hostSize failed:%w", err)
		}
		host := make([]byte, hostSize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host failed:%w", err)
		}
		ip = host
		addr = string(host)
	default:
		return errors.New("invalid atyp")
	}
	//读取端口号
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2])

	//代理转发
	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		return fmt.Errorf("dial dst failed:%w", err)
	}

	defer dest.Close()
	log.Println("dial", addr, port)

	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER socks版本，这里为0x05
	// REP Relay field,内容取值如下 X’00’ succeeded
	// RSV 保留字段
	// ATYPE 地址类型
	// BND.ADDR 服务绑定的地址
	// BND.PORT 服务绑定的端口DST.PORT

	s.successWriteHost(conn, ip, byte(Atype), port)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_, _ = io.Copy(dest, reader)
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, dest)
		cancel()
	}()

	<-ctx.Done()
	return nil
}
func (s *Socket5) successWriteHost(conn io.Writer, ip []byte, addressType byte, port uint16) error {
	//conn.Write([]byte{socks5Ver, succeeded, 0x00, atypIPV4, 0, 0, 0, 0, 0, 0})
	//return nil
	log.Printf("ip : %v,type : %v,port : %v", ip, addressType, port)
	if _, err := conn.Write([]byte{socks5Ver, succeeded, 0x00, addressType}); err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	if addressType == atypeHOST {
		length := []byte{byte(len(ip))}
		if _, err := conn.Write(length); err != nil {
			return fmt.Errorf("write failed: %w", err)
		}
	}
	if _, err := conn.Write(ip); err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	portBuf := make([]byte, 2)
	portBuf[0] = byte(port >> 8)
	portBuf[1] = byte(port - uint16(portBuf[0]))
	if _, err := conn.Write([]byte{0, 0}); err != nil {
		return fmt.Errorf("write failed: %w", err)
	}
	return nil
}
