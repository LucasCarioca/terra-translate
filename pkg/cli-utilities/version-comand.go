package cliutilities

import "fmt"

//VersionCommand used to retrieve the current version of the terraform cli
type VersionCommand struct {
	Version string
}

//Run run the command
func (c *VersionCommand) Run() error {
	fmt.Printf("Current version: v%s\n", c.Version)
	return nil
}
