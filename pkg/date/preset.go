package date

import (
	"fmt"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (formatter presetFormatter) ToTime() (time.Time, error) {
	now := time.Now()
	switch formatter.input {
	case string(NOW):
		return now, nil
	case string(TODAY):
		return time.Parse(layout, fmt.Sprintf("%d-%02d-%02dT00:00:00Z", now.Year(), now.Month(), now.Day()))
	default:
		return now, fmt.Errorf("Error: %s", formatter.input)
	}
}

func (formatter presetFormatter) ToTimeStruct(t time.Time) TimeStruct {
	return newTimeStruct(cases.Title(language.English).String(formatter.input), t)
}
