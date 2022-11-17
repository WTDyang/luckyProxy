package socks5

import "testing"

func TestSocks5Run(t *testing.T) {
	t.Run("测试socks5代理服务器", func(t *testing.T) {
		err := Run("tcp", "127.0.0.1", 1080)
		if err != nil {
			t.Fatalf("测试出现错误%e", err)
		}
	})
}
