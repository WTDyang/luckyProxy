package command

import (
	"fmt"
	"luckyProxy/client"
	"strings"
)

type (
	usageCommand struct {
	}
)

var (
	usageAll = func() {
		for _, c := range commands {
			c.help()
		}
	}
)

func (u usageCommand) help() {
	fmt.Println("  help: show usage")
	fmt.Println("        format: help [command]")
	fmt.Println("        example: help addp")
}

func (u usageCommand) run(s []string, client *client.Client) {
	var uf = usageAll

	if len(s) > 0 {
		c := commands[strings.TrimSpace(s[0])]
		if c != nil {
			uf = func() {
				c.help()
			}
		} else {
			unknownCommand(s)
			return
		}
	}
	showVersion()
	fmt.Println("有关某个命令的详细信息，请键入 help [命令名]:")
	uf()
}
