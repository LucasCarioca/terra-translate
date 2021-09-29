package terratranslate

import (
	"encoding/json"
	"strings"
)

// Changes Summary of the changes that terraform plan is showing will be made
type changes struct {
	Add       int    `json:"add"`
	Change    int    `json:"change"`
	Remove    int    `json:"remove"`
	Operation string `json:"operation"`
}

type rawChanges struct {
	Changes changes `json:"changes"`
}

func getSummary(data string) (changes, error) {
	sections := strings.SplitAfter(data, "} ")
	summaryData := sections[len(sections)-1]
	summary := rawChanges{}
	err := json.Unmarshal([]byte(summaryData), &summary)
	if err != nil {
		return changes{}, err
	}
	return summary.Changes, nil
}
