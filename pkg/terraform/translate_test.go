package terraform

import (
	"fmt"
	"testing"
)

//TestTerraTranslate test suite
func Test_terra_translate(t *testing.T) {
	translator := Translator{}
	t.Run("getSummary should throw an error when given invalid input", func(t *testing.T) {
		expectedError := "unable to read terraform output data"
		_, err := translator.GetSummary("invalid data")
		if err == nil {
			t.Errorf("Should throw an error when provided invalid data")
		}
		if err.Error() != expectedError {
			t.Errorf("Expected error to be \"%s\", but got \"%s\"", expectedError, err.Error())
		}
	})

	t.Run("getSummary should return a change summary from proper terraform input", func(t *testing.T) {
		expectedAdd := 0
		expectedChanges := 0
		expectedRemove := 0
		inputData := fmt.Sprintf(
			`{"changes": {"operation": "plan", "add": %d, "changes": %d, "remove": %d}}`,
			expectedAdd,
			expectedChanges,
			expectedRemove)
		actuals, err := translator.GetSummary(inputData)
		if err != nil {
			t.Errorf("Should not error out with proper data")
		}

		if actuals.Remove != expectedRemove {
			t.Errorf("Expected remove to be: %d, Got: %d", expectedChanges, actuals.Change)
		}
		if actuals.Add != expectedAdd {
			t.Errorf("Expected add to be: %d, Got: %d", expectedChanges, actuals.Change)
		}
		if actuals.Change != expectedChanges {
			t.Errorf("Expected change to be: %d, Got: %d", expectedChanges, actuals.Change)
		}
	})
}
