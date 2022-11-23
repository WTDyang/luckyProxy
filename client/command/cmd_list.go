package command

import (
	"fmt"
	"luckyProxy/client"
)

type (
	listCommand struct {
	}
)

func (q listCommand) help() {
	fmt.Println("  l, ls, list: show the proxy ports")
}

func (q listCommand) run(strings []string, client *client.Client) {
	if len(strings) != 0 {
		errorMsg("暂无代理端口 使用ap指令进行端口添加")
		return
	}

	infos := client.GetProxyList()

	fmt.Println("代理端口：")
	for key, proxyInfo := range infos {
		infoMsg(fmt.Sprintf(fmt.Sprintf("%s -> %s", key, proxyInfo)))
	}
}
