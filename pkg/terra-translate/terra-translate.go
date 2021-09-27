package terra_translate

import (
	"encoding/json"
	"strings"
)

type Changes struct {
	Add int				`json:"add"`
	Change int			`json:"change"`
	Remove int			`json:"remove"`
	Operation string 	`json:"operation"`
}

type rawChanges struct {
	Changes Changes `json:"changes"`
}

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