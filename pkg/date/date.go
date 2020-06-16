package date

import (
	"fmt"
	"time"
)

func (formatter dateInputFormatter) ToTime() (time.Time, error) {
	now := time.Now()
	year := now.Year()
	t, e := time.Parse(layout, fmt.Sprintf("%d-%s", year, formatter.input))
	if e == nil {
		return t, e
	}
	return time.Parse(layout, fmt.Sprintf("%d-%sT00:00:00Z", year, formatter.input[:5]))
}

func (formatter dateInputFormatter) ToTimeStruct(t time.Time) TimeStruct {
	return newTimeStruct(formatter.input, t)
}
