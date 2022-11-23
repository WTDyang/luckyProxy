package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/zeromicro/go-zero/core/conf"
	zlog "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/zerologx"
	"io/ioutil"
	utils "luckyProxy"
	"luckyProxy/client"
	"luckyProxy/client/command"
	"luckyProxy/client/handler"
	"luckyProxy/common/logx"
	"luckyProxy/common/wsx"
	"net/http"
	"os"
)

var (
	cc        = flag.String("c", "client.yaml", "配置文件路径 e.g: ./client.yaml")
	tokenFlag = flag.String("t", "", "指定连接token")
	logFile   = flag.String("l", "./client.log", "日志输出路径, e.g: ./client.log")
	username  = flag.String("u", "", "用户名, 	e.g: root")
	password  = flag.String("p", "", "密码,	e.g: ******")
	cConfig   client.Config
	token     string
)

func init() {
	//首先读取命令行参数
	flag.Parse()
	//加载配置文件
	conf.MustLoad(*cc, &cConfig)

	if *logFile != "" {
		out, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logx.Fatal().Err(err).Msg("文件打开失败")
		}
		logx.InitLogger(out)
	} else {
		logx.InitLogger(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2022/11/12 - 15:04:05"})
	}

	logx.UseLogLevel(logx.GetLogLevel(cConfig.LogLevel))
	zlog.SetWriter(zerologx.NewZeroLogWriter(logx.GetLog()))

	serverAddr := utils.FormatAddr(cConfig.Server.Host, cConfig.Server.Port)
	logx.Info().Msgf("服务器地址 %s", serverAddr)

	if *tokenFlag == utils.EmptyStr {
		token = generateToken(serverAddr)
	} else {
		token = *tokenFlag
	}

	logx.Info().Msgf("token: %s", token)
}

func main() {
	//创建连接
	c := client.NewClient(token, cConfig)

	//通过websocket进行连接
	c.Connect(func(wsx *wsx.Wsx) {
		wsx.MountBinaryFunc(handler.Dispatch(c))
	})

	//读取数据
	c.ReaderCommand(command.Dispatch)
}

//获取token
func generateToken(serverAddr string) string {
	logx.Info().Msg("未指定token，等待服务器生成......")
	mes := "username:" + *username + "-password:" + *password
	logx.Info().Msg(mes)
	url := "http://" + serverAddr + "/user/auth"
	if username != nil && *username != "" {
		url = url + "?username=" + *username + "&password=" + *password
	}
	response, err := http.Get(url)
	if err != nil {
		logx.Fatal().Err(err).Msgf("请求服务器[ %s ]生成token", serverAddr)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logx.Fatal().Err(err).Msg("服务器读取失败")
	}

	logx.Info().Msg("token获取成功")
	return string(body)
}
