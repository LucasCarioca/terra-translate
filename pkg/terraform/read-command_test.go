package terraform

import (
	"bytes"
	"errors"
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
	t.Run("Should throw and error when the command is not run with a pipe",
		func(t *testing.T) {
			expectedErrorMessage := "the command is intended to work with pipes"
			cmd := NewReadCommandWithMocks("", errors.New(expectedErrorMessage), Changes{}, nil)
			err := cmd.Run()
			assert.NotNil(t, err, "should return an error")
			assert.Equalf(t, expectedErrorMessage, err.Error(), "Should return the correct error")
		})

	t.Run("Should throw an error when the terraform output is not valid",
		func(t *testing.T) {
			expectedErrorMessage := "could not parse terraform output"
			cmd := NewReadCommandWithMocks("", nil, Changes{}, errors.New(expectedErrorMessage))
			err := cmd.Run()
			assert.NotNil(t, err, "should return an error")
			assert.Equalf(t, expectedErrorMessage, err.Error(), "Should return the correct error")
		})

	t.Run("Should read the output from terraform and display it",
		func(t *testing.T) {
			out = new(bytes.Buffer)
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
			assert.Nil(t, err, "Should not throw and error")
			assert.Equalf(t, expectedMessage, out.(*bytes.Buffer).String(), "Should return the appropriate message")
		})
}
