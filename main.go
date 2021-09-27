package main

import (
	"fmt"
	terra_translate "github.com/LucasCarioca/terra-translate/pkg/terra-translate"
	"github.com/LucasCarioca/terra-translate/pkg/utilities"
	"os"
)

func main() {
	input, err := utilities.ReadPipe()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//fmt.Printf(input)
	summary, err := terra_translate.GetSummary(input)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("operation: %s\n", summary.Operation)
	fmt.Printf("changes: %d\n", summary.Change)
	fmt.Printf("add: %d\n", summary.Add)
	fmt.Printf("destroy: %d\n", summary.Remove)
}