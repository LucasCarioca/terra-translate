package terraform

import "fmt"

//HelpCommand is used to show usage information about the cli
type HelpCommand struct{}

//Run executes the help command
func (c *HelpCommand) Run() {
	fmt.Println("USAGE: terraform [COMMAND] [OPTIONS]")
	fmt.Println("COMMANDS:")
	fmt.Println("\tread\tRead terraform logs")
	fmt.Println("\tguard\tAbort based on certain criteria")
	fmt.Println("\tversion\tGet current version")
	fmt.Println("\thelp\tGet cli documentation version")
}
