package proxy

import "testing"

func TestHandleUdp(t *testing.T) {
	t.Run("发包测试", func(t *testing.T) {
		err := udpSend("localhost:8080", []byte("hello server, i am the test"))
		if err != nil {
			t.Fatalf("出现错误%v", err)
		}
	})
}
