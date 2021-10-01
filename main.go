package main

import (
	"fmt"
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
	t "github.com/LucasCarioca/terra-translate/pkg/terraform"
	"os"
)

var (
	version string
)

func main() {
	var cmd cli.Command
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "read":
			cmd = t.NewReadCommand()
		case "guard":
			cmd = t.NewGuardCommand()
		case "version":
			cmd = &cli.VersionCommand{Version: version}
		case "help":
			cmd = &t.HelpCommand{}
		default:
			cmd = &t.HelpCommand{}
		}
	} else {
		cmd = &t.HelpCommand{}
	}
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
