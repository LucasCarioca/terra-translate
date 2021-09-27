package utilities

import (
	"bufio"
	"errors"
	"io"
	"os"
)

func ReadPipe() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return "", errors.New("The command is intended to work with pipes.\nUsage: fortune | terra-translate")
	}

	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	
	return string(output), nil
}
