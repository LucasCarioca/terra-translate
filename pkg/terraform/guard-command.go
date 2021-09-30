package terraform

import (
	"flag"
	"fmt"
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
	"os"
)

type guardOptions struct {
	destroy bool
	add bool
	change bool
}

//GuardCommand controller for reading and interpreting the terraform logs
type GuardCommand struct{
	t Translator
}

//NewGuardCommand creates a new instance of the GuardCommand
func NewGuardCommand() *GuardCommand {
	return &GuardCommand{
		t: Translator{},
	}
}

func (c *GuardCommand) getOptions() (*guardOptions, error) {
	guardCmd := flag.NewFlagSet("guard", flag.ExitOnError)
	destroy := guardCmd.Bool("d", false, "Abort when destructive changes are detected")
	add := guardCmd.Bool("a", false, "Abort when additional resource(s) are detected")
	change := guardCmd.Bool("c", false, "Abort when change(s) are detected")
	err := guardCmd.Parse(os.Args[2:])
	options := guardOptions{
		destroy: *destroy,
		add: *add,
		change: *change,
	}
	return &options, err
}

//Run executes the command
func (c *GuardCommand) Run() {

	options, err := c.getOptions()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if options.destroy || options.add || options.change {
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

		abort := false

		if options.destroy {
			if summary.Remove > 0 {
				fmt.Printf("💣 ERROR: %d destructive change(s) detected!\n", summary.Remove)
				abort = true
			} else {
				fmt.Println("🚀 No destructive changes detected")
			}
		}

		if options.add {
			if summary.Add > 0 {
				fmt.Printf("💣 ERROR: %d additional resource(s) detected!\n", summary.Add)
				abort = true
			} else {
				fmt.Println("🚀 No additional resources detected")
			}
		}

		if options.change {
			if summary.Change > 0 {
				fmt.Printf("💣 ERROR: %d resource change(s) detected!\n", summary.Change)
				abort = true
			} else {
				fmt.Println("🚀 No resources to be changed detected")
			}
		}

		if abort {
			fmt.Println("\n\nGuarded changes have been detected.\n" +
				"See the output above for more information.\n" +
				"Exiting with code 1")
			os.Exit(1)
		}
	}
}
