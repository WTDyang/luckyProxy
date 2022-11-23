package command

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/color"
	"luckyProxy/client"
	"strings"
)

type command interface {
	help()
	run([]string, *client.Client)
}

var (
	commands = map[string]command{}
	version  = "lucky::1.2.4"
)

func init() {
	commands["help"] = &usageCommand{}
	commands["quit"] = &quitCommand{}
	commands["addp"] = &addProxyCommand{}
	commands["removep"] = &removeProxyCommand{}
	commands["list"] = &listCommand{}
}

func Dispatch(line string, client *client.Client) {
	split := strings.Split(line, " ")
	switch strings.TrimSpace(split[0]) {
	case "usage", "h", "help", "?":
		commands["help"].run(split[1:], client)
	case "q", "quit", "exit":
		commands["quit"].run(split, client)
	case "ap", "add-port", "addp":
		commands["addp"].run(split[1:], client)
	case "removep", "rp":
		commands["removep"].run(split[1:], client)
	case "list", "ls", "l":
		commands["list"].run(split[1:], client)
	case "version":
		showVersion()
	default:
		unknownCommand(split)
	}
}

func showVersion() {
	fmt.Printf("lucky-poxy %s\n", version)
}

func unknownCommand(cmd []string) {
	errorMsg("unknown command: " + cmd[0])
}

func errorMsg(msg string) {
	print(msg, color.BgRed)
}

func infoMsg(msg string) {
	print(msg, color.FgGreen)
}

func print(msg string, colour color.Color) {
	fmt.Println(color.WithColor("[cli]", color.FgBlue)+":", color.WithColor(msg, colour))
}
