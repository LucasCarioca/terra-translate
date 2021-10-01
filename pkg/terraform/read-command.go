package terraform

import (
	"flag"
	"fmt"
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
	"os"
)

type readOptions struct{}

//ReadCommand controller for reading and interpreting the terraform logs
type ReadCommand struct {
	t    Translator
	pipe func() (string, error)
}

//NewReadCommand creates a new instance of the ReadCommand
func NewReadCommand() *ReadCommand {
	return &ReadCommand{
		t:    Translator{},
		pipe: cli.ReadPipe,
	}
}

func (_ *ReadCommand) getOptions() (*readOptions, error) {
	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	err := readCmd.Parse(os.Args[2:])
	options := readOptions{}
	return &options, err
}

//Run executes the command
func (c *ReadCommand) Run() error {

	_, err := c.getOptions()
	if err != nil {
		return err
	}

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
