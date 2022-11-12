package command

import (
	"fmt"
	"luckyProxy/client"
	"os"
)

type (
	quitCommand struct {
	}
)

func (q quitCommand) help() {
	fmt.Println("  q, quit, exit: quit client")
}

func (q quitCommand) run(strings []string, client *client.Client) {
	client.Close()
	infoMsg("感谢使用luckyProxy~\n有缘再见!")
	os.Exit(0)
}
