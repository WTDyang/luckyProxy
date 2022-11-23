package nosqlDB

import (
	"strings"
	"testing"
)

func TestGetConn(t *testing.T) {
	t.Run("获取redis客户端", func(t *testing.T) {
		conn := GetConn()
		if conn == nil {
			t.Fatalf("测试失败%s", "客户端本该存在但是为空")
		}
		pong, err := conn.Ping().Result()
		if err != nil {
			t.Fatalf("测试失败,出现不必要错误%e", err)
		}
		if !strings.EqualFold(pong, "pong") {
			t.Fatalf("测试失败,本应该出现pong ,却为%s", pong)
		}
	})
}
