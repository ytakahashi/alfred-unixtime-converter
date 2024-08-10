package date

import (
	"fmt"
	"strconv"
	"time"
)

func (formatter unixTimeInputFormatter) ToTime() (time.Time, error) {
	unixtime, _ := strconv.ParseInt(formatter.input, 10, 64)
	if formatter.unit == "ms" {
		t := time.UnixMilli(unixtime)
		return t, nil
	} else {
		t := time.Unix(unixtime, 0)
		return t, nil
	}
}

func (formatter unixTimeInputFormatter) ToTimeStruct(t time.Time) TimeStruct {
	return newTimeStruct(fmt.Sprintf("%s (%s)", formatter.input, formatter.unit), t)
}
