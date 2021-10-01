package terraform

type mockTranslator struct {
	mockSummaryChanges Changes
	mockSummaryError   error
}

//GetSummary mocked version of the get summary method
func (t *mockTranslator) GetSummary(_ string) (Changes, error) {
	return t.mockSummaryChanges, t.mockSummaryError
}