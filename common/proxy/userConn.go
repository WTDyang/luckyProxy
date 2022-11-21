package proxy

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"luckyProxy/common/logx"
	"luckyProxy/common/protocal"
	"net"
	"strings"
)

type (
	UserConn struct {
		// Id uuid
		Id   string
		conn net.Conn
		c    *Container
		// key used to identify the service started by the server
		key       string
		writeChan chan []byte
	}
)

func NewUserConn(conn net.Conn, c *Container, key string) *UserConn {
	return &UserConn{
		Id:        uuid.New().String(),
		conn:      conn,
		c:         c,
		key:       key,
		writeChan: make(chan []byte),
	}
}

// OnUserConnect 通知客户端进行客户端监听
func (u UserConn) OnUserConnect() error {
	bytes, err := protocal.NewUserConnect(u.key, u.Id).Encode()
	if err != nil {
		u.err(err).Msg("encode userConnect")
		return err
	}

	u.c.WriteBinary(bytes)
	return nil
}

// StartRead 读取用户对客户端的请求以写入intranet服务 .
func (u UserConn) StartRead(clean func()) {
	defer clean()

	for {
		// todo read buffer size
		buf := make([]byte, 1024)
		n, err := u.conn.Read(buf)
		if err != nil {
			if strings.ContainsAny("use of closed network connection", err.Error()) || strings.ContainsAny("EOF", err.Error()) {
				return
			}

			u.err(err).Msg("read user connection")
			return
		}

		if n == 0 {
			continue
		}

		bytes, err := protocal.NewUserRequest(buf[:n], u.key, u.Id).Encode()

		if err != nil {
			u.err(err).Msg("encode userRequest")
			continue
		}

		u.c.WriteBinary(bytes)
		logx.Debug().Int("write", n).Msg("write to client on user request")
	}
}

// StartWrite 开始向用户写入intranet响应
func (u UserConn) StartWrite(clean func()) {
	defer clean()

	for {
		select {
		case data := <-u.writeChan:
			// write user request data to internet
			_, err := u.conn.Write(data)
			if err != nil {
				u.err(err).Msg("write to user")
				return
			}
		}
	}
}

// Write data to user
func (u UserConn) Write(data []byte) {
	u.writeChan <- data
}

func (u UserConn) err(err error) *zerolog.Event {
	return logx.Err(err).Str("connId", u.Id).Str("key", u.key)
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
