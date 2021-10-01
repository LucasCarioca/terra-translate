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

func Test_guard_command(t *testing.T) {
	t.Run("Should throw error when input is not piped to command",
		func(t *testing.T) {
			cli.MockCLICall("terra-translate guard -d")
			expectedErrorMessage := "the command is intended to work with pipes"
			cmd := NewGuardCommandWithMocks("", errors.New(expectedErrorMessage), Changes{}, nil)
			err := cmd.Run()
			assert.NotNil(t, err, "should return an error")
			assert.Equalf(t, expectedErrorMessage, err.Error(), "should return the correct error")
		})

	t.Run("Should throw error when piped a valid terraform output",
		func(t *testing.T) {
			cli.MockCLICall("terra-translate guard -d")
			expectedErrorMessage := "could not parse terraform output"
			cmd := NewGuardCommandWithMocks("", nil, Changes{}, errors.New(expectedErrorMessage))
			err := cmd.Run()
			assert.NotNil(t, err, "should return an error")
			assert.Equalf(t, expectedErrorMessage, err.Error(), "should return the correct error")
		})

	t.Run("Should return no destructive changes when run with the destructive guard option and no destructive changes are found",
		func(t *testing.T) {
			out = new(bytes.Buffer)
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

	t.Run("Should abort when run with the destructive guard option and destructive changes are found",
		func(t *testing.T) {
			out = new(bytes.Buffer)
			cli.MockCLICall("terra-translate guard -d")
			expectedMessage := "💣 ERROR: 1 destructive change(s) detected!\n"
			expectedErrorMessage := "\n\nGuarded changes have been detected.\n" +
				"See the output above for more information.\n" +
				"Exiting with code 1"
			expectedChanges := Changes{
				Operation: "plan",
				Add:       0,
				Change:    0,
				Remove:    1,
			}
			cmd := NewGuardCommandWithMocks("", nil, expectedChanges, nil)
			err := cmd.Run()
			assert.NotNil(t, err, "should throw an error")
			assert.Equalf(t, expectedMessage, out.(*bytes.Buffer).String(), "should return the appropriate message")
			assert.Equalf(t, expectedErrorMessage, err.Error(), "should throw the appropriate error")
		})

	t.Run("Should return no additional resources when run with the add guard option and no additional resources are found",
		func(t *testing.T) {
			out = new(bytes.Buffer)
			cli.MockCLICall("terra-translate guard -a")
			expectedMessage := "🚀 No additional resources detected\n"
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

	t.Run("Should abort when run with the add guard option and additional resources are found",
		func(t *testing.T) {
			out = new(bytes.Buffer)
			cli.MockCLICall("terra-translate guard -a")
			expectedMessage := "💣 ERROR: 1 additional resource(s) detected!\n"
			expectedErrorMessage := "\n\nGuarded changes have been detected.\n" +
				"See the output above for more information.\n" +
				"Exiting with code 1"
			expectedChanges := Changes{
				Operation: "plan",
				Add:       1,
				Change:    0,
				Remove:    0,
			}
			cmd := NewGuardCommandWithMocks("", nil, expectedChanges, nil)
			err := cmd.Run()
			assert.NotNil(t, err, "should throw an error")
			assert.Equalf(t, expectedMessage, out.(*bytes.Buffer).String(), "should return the appropriate message")
			assert.Equalf(t, expectedErrorMessage, err.Error(), "should throw the appropriate error")
		})

	t.Run("Should return no changed resources when run with the change guard option and no changed resources are found",
		func(t *testing.T) {
			out = new(bytes.Buffer)
			cli.MockCLICall("terra-translate guard -c")
			expectedMessage := "🚀 No resources to be changed detected\n"
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

	t.Run("Should abort when run with the change guard option and changed resources are found",
		func(t *testing.T) {
			out = new(bytes.Buffer)
			cli.MockCLICall("terra-translate guard -c")
			expectedMessage := "💣 ERROR: 1 resource change(s) detected!\n"
			expectedErrorMessage := "\n\nGuarded changes have been detected.\n" +
				"See the output above for more information.\n" +
				"Exiting with code 1"
			expectedChanges := Changes{
				Operation: "plan",
				Add:       0,
				Change:    1,
				Remove:    0,
			}
			cmd := NewGuardCommandWithMocks("", nil, expectedChanges, nil)
			err := cmd.Run()
			assert.NotNil(t, err, "should throw an error")
			assert.Equalf(t, expectedMessage, out.(*bytes.Buffer).String(), "should return the appropriate message")
			assert.Equalf(t, expectedErrorMessage, err.Error(), "should throw the appropriate error")
		})
}
