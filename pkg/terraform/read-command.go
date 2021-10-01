package terraform

import (
	"fmt"
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
)

type readOptions struct{}

//ReadCommand controller for reading and interpreting the terraform logs
type ReadCommand struct {
	t    TranslatorInterface
	pipe func() (string, error)
}

//NewReadCommand creates a new instance of the ReadCommand
func NewReadCommand() *ReadCommand {
	return &ReadCommand{
		t:    &Translator{},
		pipe: cli.ReadPipe,
	}
}

//Run executes the command
func (c *ReadCommand) Run() error {

	input, err := c.pipe()
	if err != nil {
		return err
	}

	summary, err := c.t.GetSummary(input)
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "operation: %s\n", summary.Operation)
	fmt.Fprintf(out, "changes: %d\n", summary.Change)
	fmt.Fprintf(out, "add: %d\n", summary.Add)
	fmt.Fprintf(out, "destroy: %d\n", summary.Remove)

	return nil
}
