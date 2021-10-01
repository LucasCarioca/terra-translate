package terraform

import (
	"bytes"
	"errors"
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewReadCommandWithMocks(pipeContent string, pipeError error, changesContent Changes, changesError error) *ReadCommand {
	return &ReadCommand{
		t: &mockTranslator{
			mockSummaryChanges: changesContent,
			mockSummaryError:   changesError,
		},
		pipe: func() (string, error) {
			return pipeContent, pipeError
		},
	}
}

func Test_read_command(t *testing.T) {
	t.Run("", func(t *testing.T) {
		cli.MockCLICall("terra-translate guard -d")
		expectedErrorMessage := "the command is intended to work with pipes"
		cmd := NewReadCommandWithMocks("", errors.New(expectedErrorMessage), Changes{}, nil)
		err := cmd.Run()
		assert.NotNil(t, err, "should return an error")
		assert.Equalf(t, expectedErrorMessage, err.Error(), "should return the correct error")
	})

	t.Run("", func(t *testing.T) {
		cli.MockCLICall("terra-translate guard -d")
		expectedErrorMessage := "could not parse terraform output"
		cmd := NewReadCommandWithMocks("", nil, Changes{}, errors.New(expectedErrorMessage))
		err := cmd.Run()
		assert.NotNil(t, err, "should return an error")
		assert.Equalf(t, expectedErrorMessage, err.Error(), "should return the correct error")
	})

	t.Run("", func(t *testing.T) {
		out = new(bytes.Buffer)
		cli.MockCLICall("terra-translate guard -d")
		expectedMessage := "operation: plan\n" +
			"changes: 0\n" +
			"add: 0\n" +
			"destroy: 0\n"
		expectedChanges := Changes{
			Operation: "plan",
			Add:       0,
			Change:    0,
			Remove:    0,
		}
		cmd := NewReadCommandWithMocks("", nil, expectedChanges, nil)
		err := cmd.Run()
		assert.Nil(t, err, "should not throw and error")
		assert.Equalf(t, expectedMessage, out.(*bytes.Buffer).String(), "should return the appropriate message")
	})
}
