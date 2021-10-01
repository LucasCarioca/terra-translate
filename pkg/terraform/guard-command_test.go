package terraform

import (
	"bytes"
	"errors"
	cli "github.com/LucasCarioca/terra-translate/pkg/cli-utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewGuardCommandWithMocks(pipeContent string, pipeError error, changesContent Changes, changesError error) *GuardCommand {
	return &GuardCommand{
		t: &mockTranslator{
			mockSummaryChanges: changesContent,
			mockSummaryError:   changesError,
		},
		pipe: func() (string, error) {
			return pipeContent, pipeError
		},
	}
}

type mockTranslator struct {
	mockSummaryChanges Changes
	mockSummaryError   error
}

//GetSummary mocked version of the get summary method
func (t *mockTranslator) GetSummary(data string) (Changes, error) {
	return t.mockSummaryChanges, t.mockSummaryError
}

func Test_guard_command(t *testing.T) {
	out = new(bytes.Buffer)

	t.Run("Should throw error when input is not piped to command",
		func(t *testing.T) {
			cli.MockCLICall("terra-translate guard -d")
			expectedErrorMessage := "the command is intended to work with pipes"
			cmd := NewGuardCommandWithMocks("", errors.New(expectedErrorMessage), Changes{}, nil)
			err := cmd.Run()
			assert.NotNil(t, err, "should return an error")
			assert.Equalf(t, expectedErrorMessage, err.Error(), "should return the correct error")
		})

	t.Run("Should return no destructive changes when run with the destructive guard option and no destructive changes are found",
		func(t *testing.T) {
			cli.MockCLICall("terra-translate guard -d")
			expectedMessage := "🚀 No destructive changes detected\n"
			expectedChanges := Changes{
				Operation: "plan",
				Add:       0,
				Change:    0,
				Remove:    0,
			}
			cmd := NewGuardCommandWithMocks("", nil, expectedChanges, nil)
			err := cmd.Run()
			assert.Nil(t, err, "should not throw and error")
			assert.Equalf(t, expectedMessage, out.(*bytes.Buffer).String(), "should return the appropriate message")
		})
}
