package main

import (
	"flag"
	"fmt"
	t "github.com/LucasCarioca/terra-translate/pkg/terra-translate"
	"github.com/LucasCarioca/terra-translate/pkg/utilities"
	"os"
)

func main() {
	destroyGuard := flag.Bool("destroy-guard", false, "A bool: when true, exit code will be 1 when destructive changes are detected.")
	flag.Parse()
	input, err := utilities.ReadPipe()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//fmt.Printf(input)
	summary, err := t.GetSummary(input)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("operation: %s\n", summary.Operation)
	fmt.Printf("changes: %d\n", summary.Change)
	fmt.Printf("add: %d\n", summary.Add)
	fmt.Printf("destroy: %d\n", summary.Remove)

	if *destroyGuard && summary.Remove > 0 {
		fmt.Printf("WARNING: %d destructive change(s) detected!\n", summary.Remove)
		os.Exit(1)
	}
}