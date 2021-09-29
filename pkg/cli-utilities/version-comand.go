package cliutilities

import "fmt"

//VersionCommand used to retrieve the current version of the terra-translate cli
type VersionCommand struct {
	Version string
}

//Run run the command
func (c *VersionCommand) Run() {
	fmt.Printf("Current version: v%s\n", c.Version)
}
