package terraform

import (
	"errors"
	"flag"
	"fmt"
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
	"os"
)

type guardOptions struct {
	destroy bool
	add     bool
	change  bool
}

//GuardCommand controller for reading and interpreting the terraform logs
type GuardCommand struct {
	t    TranslatorInterface
	pipe func() (string, error)
}

//NewGuardCommand creates a new instance of the GuardCommand
func NewGuardCommand() *GuardCommand {
	return &GuardCommand{
		t:    &Translator{},
		pipe: cli.ReadPipe,
	}
}

func (*GuardCommand) getOptions() (*guardOptions, error) {
	guardCmd := flag.NewFlagSet("guard", flag.ExitOnError)
	destroy := guardCmd.Bool("d", false, "Abort when destructive changes are detected")
	add := guardCmd.Bool("a", false, "Abort when additional resource(s) are detected")
	change := guardCmd.Bool("c", false, "Abort when change(s) are detected")
	err := guardCmd.Parse(os.Args[2:])
	options := guardOptions{
		destroy: *destroy,
		add:     *add,
		change:  *change,
	}
	return &options, err
}

//Run executes the command
func (c *GuardCommand) Run() error {

	options, err := c.getOptions()
	if err != nil {
		return err
	}

	if options.destroy || options.add || options.change {
		input, err := c.pipe()
		if err != nil {
			return err
		}

		summary, err := c.t.GetSummary(input)
		if err != nil {
			return err
		}

		abort := false

		if options.destroy {
			if summary.Remove > 0 {
				fmt.Fprintf(out, "💣 ERROR: %d destructive change(s) detected!\n", summary.Remove)
				abort = true
			} else {
				fmt.Fprintln(out, "🚀 No destructive changes detected")
			}
		}

		if options.add {
			if summary.Add > 0 {
				fmt.Fprintf(out, "💣 ERROR: %d additional resource(s) detected!\n", summary.Add)
				abort = true
			} else {
				fmt.Fprintln(out, "🚀 No additional resources detected")
			}
		}

		if options.change {
			if summary.Change > 0 {
				fmt.Fprintf(out, "💣 ERROR: %d resource change(s) detected!\n", summary.Change)
				abort = true
			} else {
				fmt.Fprintln(out, "🚀 No resources to be changed detected")
			}
		}

		if abort {
			return errors.New("\n\nGuarded changes have been detected.\n" +
				"See the output above for more information.\n" +
				"Exiting with code 1")
		}
	}
	return nil
}
