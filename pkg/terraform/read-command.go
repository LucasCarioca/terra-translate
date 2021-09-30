package terraform

import (
	"fmt"
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
	"os"
)

//ReadCommand controller for reading and interpreting the terraform logs
type ReadCommand struct {
	t Translator
}

//NewReadCommand creates a new instance of the ReadCommand
func NewReadCommand() *ReadCommand {
	return &ReadCommand{
		t: Translator{},
	}
}

//Run executes the command
func (c *ReadCommand) Run() {
	input, err := cli.ReadPipe()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	summary, err := c.t.GetSummary(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("operation: %s\n", summary.Operation)
	fmt.Printf("changes: %d\n", summary.Change)
	fmt.Printf("add: %d\n", summary.Add)
	fmt.Printf("destroy: %d\n", summary.Remove)
}
