package main

import (
	"fmt"
	"github.com/LucasCarioca/terra-translate/pkg/utilities"
	"os"
)

func main() {
	input, err := utilities.ReadPipe()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf(input)
}