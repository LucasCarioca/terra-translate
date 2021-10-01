package terraform

import "fmt"

//HelpCommand is used to show usage information about the cli
type HelpCommand struct{}

//Run executes the help command
func (_ *HelpCommand) Run() error {
	fmt.Fprintln(out, "USAGE: terraform [COMMAND] [OPTIONS]")
	fmt.Fprintln(out, "COMMANDS:")
	fmt.Fprintln(out, "\tread\tRead terraform logs")
	fmt.Fprintln(out, "\tguard\tAbort based on certain criteria")
	fmt.Fprintln(out, "\tversion\tGet current version")
	fmt.Fprintln(out, "\thelp\tGet cli documentation version")
	return nil
}
