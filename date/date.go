package date

import (
	"strconv"
	"time"
)

// TimeStruct is a struct
type TimeStruct struct {
	Unixtime       int64
	UnixtimeMillis int64
	DateTime       string
}

var layout = time.RFC3339

// NewTimeStruct creates a new TimeStruct from string value
func NewTimeStruct(value string) TimeStruct {
	t, e := fromDateTimeString(value)
	if e == nil {
		return newTimeStruct(t)
	}

	t2, e2 := fromUnixTimestamp(value)
	if e2 == nil {
		return newTimeStruct(t2)
	}

	now := time.Now()
	return newTimeStruct(now)
}

func newTimeStruct(t time.Time) TimeStruct {
	return TimeStruct{
		Unixtime:       t.Unix(),
		UnixtimeMillis: t.UnixNano() / int64(time.Millisecond),
		DateTime:       t.UTC().Format(layout),
	}
}

func fromDateTimeString(value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func fromUnixTimestamp(value string) (time.Time, error) {
	unixtime, e := strconv.ParseInt(value, 10, 64)
	if e != nil {
		return time.Time{}, e
	}

	t := time.Unix(unixtime, 0)
	return t, nil
}
