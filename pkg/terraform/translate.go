package terraform

import (
	"encoding/json"
	"errors"
	"strings"
)

//TranslatorInterface interface for translator
type TranslatorInterface interface {
	GetSummary(data string) (Changes, error)
}

//Translator a utility to interpret the output from terraform operations
type Translator struct{}

//Changes Summary of the changes that terraform plan is showing will be made
type Changes struct {
	Add       int    `json:"add"`
	Change    int    `json:"change"`
	Remove    int    `json:"remove"`
	Operation string `json:"operation"`
}

//ChangesWrapper wraps the changes data and contains other metadata
type ChangesWrapper struct {
	Changes Changes `json:"changes"`
}

//GetSummary takes the Terraform output as a parameter and returns a summary of the changes
func (*Translator) GetSummary(data string) (Changes, error) {
	sections := strings.SplitAfter(data, "} ")
	summaryData := sections[len(sections)-1]
	summary := ChangesWrapper{}
	err := json.Unmarshal([]byte(summaryData), &summary)
	if err != nil {
		return Changes{}, errors.New("unable to read terraform output data")
	}
	return summary.Changes, nil
}
