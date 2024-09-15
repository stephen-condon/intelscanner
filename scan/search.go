package scan

import (
	"fmt"
	"regexp"
)

type SearchConfig struct {
	Data []*ParsedLine
}

func NewSearcher(data []*ParsedLine) *SearchConfig {
	return &SearchConfig{
		Data: data,
	}
}

// for now, just search base
func (sc *SearchConfig) SearchBase(str string) []*ParsedLine {
	var result []*ParsedLine
	searchString := fmt.Sprintf(`.*%v.*`, str)
	re := regexp.MustCompile(searchString)
	for _, record := range sc.Data {
		match := re.FindString(record.Base)
		if len(match) != 0 {
			result = append(result, record)
		}
	}

	return result
}
