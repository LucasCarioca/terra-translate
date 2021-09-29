package main

import (
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
	t "github.com/LucasCarioca/terra-translate/pkg/terra-translate"
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
			cmd = &t.ReadCommand{}
		case "guard":
			cmd = &t.GuardCommand{}
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
	cmd.Run()
}
