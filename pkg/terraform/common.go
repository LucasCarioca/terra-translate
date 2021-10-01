package terraform

import (
	"io"
	"os"
)

var (
	out  io.Writer      = os.Stdout
)
