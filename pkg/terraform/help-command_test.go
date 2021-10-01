package terraform

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_help_command(t *testing.T) {
	t.Run("Should display the usage information for the cli",
		func(t *testing.T) {
			out = new(bytes.Buffer)
			expectedContent := "USAGE: terraform [COMMAND] [OPTIONS]\n" +
				"COMMANDS:\n" +
				"\tread\tRead terraform logs\n" +
				"\tguard\tAbort based on certain criteria\n" +
				"\tversion\tGet current version\n" +
				"\thelp\tGet cli documentation version\n"
			cmd := HelpCommand{}
			err := cmd.Run()
			assert.Nil(t, err, "should not throw an error")
			assert.Equal(t, out.(*bytes.Buffer).String(), expectedContent, "Should return the expected text")
		})
}
