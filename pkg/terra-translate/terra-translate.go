package terratranslate

import (
	"encoding/json"
	"strings"
)

// Changes Summary of the changes that terraform plan is showing will be made
type Changes struct {
	Add int				`json:"add"`
	Change int			`json:"change"`
	Remove int			`json:"remove"`
	Operation string 	`json:"operation"`
}

type rawChanges struct {
	Changes Changes `json:"changes"`
}

// GetSummary returns a set of changes from the provided terraform plan
func GetSummary(data string) (Changes, error) {
	sections := strings.SplitAfter(data, "} ")
	summaryData := sections[len(sections)-1]
	summary := rawChanges{}
	err := json.Unmarshal([]byte(summaryData), &summary)
	if err != nil {
		return Changes{}, err
	}
	return summary.Changes, nil
}