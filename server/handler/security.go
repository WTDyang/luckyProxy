package handler

import (
	"fmt"
	zlog "github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"luckyProxy/server/nosqlDB"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) Register() bool {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost) //盐值加密处理
	if err != nil {
		fmt.Println(err)
		return false
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println(encodePWD)
	conn := nosqlDB.GetConn()
	conn.Set("user:"+u.Name, encodePWD, 0).Err()
	return true
}
func (u *User) Login() bool {
	//encodePWD := "$2a$10$ZrMynkc1fkDcRhuxHd./gujzPpm.ImFyhYe2CYMxidJponCYQSOmm" //模拟从数据库中读取到的 经过bcrypt.GenerateFromPassword处理的密码值
	conn := nosqlDB.GetConn()
	encodePWD, err := conn.Get("user:" + u.Name).Result()
	if err != nil {
		zlog.Error(err)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(encodePWD), []byte(u.Password)) //验证（对比）
	if err != nil {
		fmt.Println("pwd wrong")
		return false
	} else {
		fmt.Println("pwd ok")
		return true
	}
}
