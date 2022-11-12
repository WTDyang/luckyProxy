package command

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	utils "luckyProxy"
	"luckyProxy/client"
	pkg "luckyProxy/common"
	"luckyProxy/common/model/req"
	"net/url"
	"strings"
)

type (
	addProxyCommand struct{}
)

func (a addProxyCommand) help() {
	fmt.Println("  addp: add proxy ")
	fmt.Println("      format: addp [channelType]:[ip]:[port]")
	fmt.Println("      example:")
	fmt.Println("               addp tcp::8080")
	fmt.Println("               tcp:192.168.0.1:5555")
}

func (a addProxyCommand) callHelp() {
	Dispatch("help addp", nil)
}

func (a addProxyCommand) run(s []string, c *client.Client) {
	if len(s) == 0 {
		errorMsg("输入错误")
		return
	}

	var infos []req.AddProxyInfo

	for _, line := range s {
		split := strings.Split(line, ":")
		if len(split) < 3 {
			errorMsg("proxy format error: " + line)
			a.callHelp()
			return
		}

		port, err := cast.ToIntE(strings.TrimSuffix(split[2], "\r"))
		if err != nil {
			errorMsg(fmt.Sprintf("port %s is not valid", split[2]))
			a.callHelp()
			return
		}

		var ip string
		if split[1] == utils.EmptyStr {
			ip = "localhost"
		} else {
			ip = split[1]
		}

		proxyInfo := req.AddProxyInfo{
			ChannelType: split[0],
			Ip:          ip,
			Port:        cast.ToInt(port),
		}
		infos = append(infos, proxyInfo)
	}

	proxyInfoReq := req.AddProxyInfoReq{Proxy: infos}
	err := proxyInfoReq.Check()
	if err != nil {
		errorMsg(err.Error())
		a.callHelp()
		return
	}

	u := url.URL{Path: AddProxy + c.Token(), Scheme: "http", Host: c.ServerAddr()}

	resp, err := PostJson(u, proxyInfoReq)
	if err != nil {
		errorMsg(err.Error())
		return
	}

	f, response := ShowResp(resp)

	var proxyInfos []pkg.ClientProxyInfo
	err = json.Unmarshal(response, &proxyInfos)
	if err != nil {
		f(string(response))
		return
	}

	for _, proxyInfo := range proxyInfos {
		f("add proxy:")
		//TODO 地址有问题
		f("    " + fmt.Sprintf("%s -> %s", proxyInfo.IntranetAddr, proxyInfo.Address(utils.GetCurrentIp())))
	}
}
