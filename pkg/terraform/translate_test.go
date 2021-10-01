package terraform

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

//TestTerraTranslate test suite
func Test_terra_translate(t *testing.T) {
	translator := Translator{}

	t.Run("getSummary should throw an error when given invalid terraform output",
		func(t *testing.T) {
			expectedError := "unable to read terraform output data"
			_, err := translator.GetSummary("invalid data")
			assert.NotNil(t, err, "Should throw an error")
			assert.Equalf(t, expectedError, err.Error(), "Should throw the appropriate error")
		})

	t.Run("getSummary should return a change summary from proper terraform output",
		func(t *testing.T) {
			expectedAdd := rand.Int()
			expectedChanges := rand.Int()
			expectedRemove := rand.Int()
			inputData := fmt.Sprintf(
				`{"changes": {"operation": "plan", "add": %d, "change": %d, "remove": %d}}`,
				expectedAdd,
				expectedChanges,
				expectedRemove)
			actual, err := translator.GetSummary(inputData)
			assert.Nil(t, err, "Should not throw an error")
			assert.Equalf(t, expectedRemove, actual.Remove, "Should match data from terraform input (destroy)")
			assert.Equalf(t, expectedAdd, actual.Add, "Should match data from terraform input (add)")
			assert.Equalf(t, expectedChanges, actual.Change, "Should match data from terraform input (change)")
		})

	t.Run("getSummary should return a change summary even if there are multiple rows of output from terraform",
		func(t *testing.T) {
			expectedAdd := rand.Int()
			expectedChanges := rand.Int()
			expectedRemove := rand.Int()
			inputData := fmt.Sprintf(
				`{} {} {} {"changes": {"operation": "plan", "add": %d, "change": %d, "remove": %d}}`,
				expectedAdd,
				expectedChanges,
				expectedRemove)
			_, err := translator.GetSummary(inputData)
			assert.Nil(t, err, "Should not throw an error")
		})
}
