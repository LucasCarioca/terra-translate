package terraform

import (
	"flag"
	"fmt"
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
	"os"
)

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

//Run executes the command
func (c *GuardCommand) Run() {

	guardCmd := flag.NewFlagSet("guard", flag.ExitOnError)
	destroy := guardCmd.Bool("d", false, "Abort when destructive changes are detected")
	add := guardCmd.Bool("a", false, "Abort when additional resource(s) are detected")
	change := guardCmd.Bool("c", false, "Abort when change(s) are detected")
	err := guardCmd.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if *destroy || *add || *change {
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

		if *destroy {
			if summary.Remove > 0 {
				fmt.Printf("💣 ERROR: %d destructive change(s) detected!\n", summary.Remove)
				abort = true
			} else {
				fmt.Println("🚀 No destructive changes detected")
			}
		}

		if *add {
			if summary.Add > 0 {
				fmt.Printf("💣 ERROR: %d additional resource(s) detected!\n", summary.Add)
				abort = true
			} else {
				fmt.Println("🚀 No additional resources detected")
			}
		}

		if *change {
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
	} else {
		fmt.Println("⚠️ Nothing to do. Please select ad least one option")
		guardCmd.Usage()
		os.Exit(1)
	}
}
