package handler

import "testing"

func TestRegister(t *testing.T) {
	t.Run("register success", func(t *testing.T) {
		user := User{
			Name:     "yxk",
			Password: "123",
		}
		if ok := user.Register(); !ok {
			t.Fatalf("注册测试失败:%s", "参数无误但是注册失败")
		}
	})
}
func TestLogin(t *testing.T) {
	t.Run("login success", func(t *testing.T) {
		user := User{
			Name:     "yxk",
			Password: "123",
		}
		if ok := user.Login(); !ok {
			t.Fatalf("登录测试失败%s", "账号密码正确但是测试失败\n")
		}
	})
	t.Run("login when password wrong", func(t *testing.T) {
		user := User{
			Name:     "yxk",
			Password: "123456",
		}
		if ok := user.Login(); ok {
			t.Fatalf("登录测试失败%s", "账号密码错误但是测试成功\n")
		}
	})
}
