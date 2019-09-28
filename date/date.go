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

type formatter interface {
	newTimeStruct(t time.Time) TimeStruct
	fromDateTimeString(value string) (time.Time, error)
	fromUnixTimestamp(value string) (time.Time, error)
}

var layout = time.RFC3339

// NewTimeStruct creates a new TimeStruct from string value
func NewTimeStruct(value string, formatter formatter) TimeStruct {
	t, e := formatter.fromDateTimeString(value)
	if e == nil {
		return formatter.newTimeStruct(t)
	}

	t2, e2 := formatter.fromUnixTimestamp(value)
	if e2 == nil {
		return formatter.newTimeStruct(t2)
	}

	now := time.Now()
	return formatter.newTimeStruct(now)
}

// TimeStructFormatter TimeStructFormatter
type TimeStructFormatter struct{}

func (formatter TimeStructFormatter) newTimeStruct(t time.Time) TimeStruct {
	return TimeStruct{
		Unixtime:       t.Unix(),
		UnixtimeMillis: t.UnixNano() / int64(time.Millisecond),
		DateTime:       t.UTC().Format(layout),
	}
}

func (formatter TimeStructFormatter) fromDateTimeString(value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func (formatter TimeStructFormatter) fromUnixTimestamp(value string) (time.Time, error) {
	unixtime, e := strconv.ParseInt(value, 10, 64)
	if e != nil {
		return time.Time{}, e
	}

	t := time.Unix(unixtime, 0)
	return t, nil
}
