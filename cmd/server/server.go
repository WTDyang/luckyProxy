package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/zeromicro/go-zero/core/conf"
	zlog "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/zero-contrib/logx/zerologx"
	"luckyProxy/common/logx"
	"luckyProxy/server"
	"luckyProxy/server/api"
	"luckyProxy/server/svc"
	"os"
)

var (
	sc      = flag.String("c", "server.yaml", "the config file path")
	logFile = flag.String("l", "", "the log file path, e.g: ./server.log")
	sConfig server.Config
)

func init() {
	flag.Parse()

	conf.MustLoad(*sc, &sConfig)

	if *logFile != "" {
		out, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logx.Fatal().Err(err).Msg("open log file fail")
		}
		logx.InitLogger(out)
	} else {
		logx.InitLogger(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2022/11/12 - 15:04:05"})
	}

	logx.UseLogLevel(logx.GetLogLevel(sConfig.LogLevel))
	zlog.SetWriter(zerologx.NewZeroLogWriter(logx.GetLog()))
}

func main() {

	server := rest.MustNewServer(sConfig.RestConf)
	defer server.Stop()
	svcContext := svc.NewServiceContext(sConfig)

	api.MountRouters(server, svcContext)

	server.Start()
}
