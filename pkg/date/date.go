package date

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var hhmmss = regexp.MustCompile(`^[0-2]?[0-9]:[0-5][0-9]:[0-5][0-9]`)
var mmdd = regexp.MustCompile(`^[0-1][0-9]\-[0-3][0-9]`)

// TimeStruct is a struct
type TimeStruct struct {
	Value          string
	Unixtime       int64
	UnixtimeMillis int64
	DateTime       string
	LocalDateTime  string
}

type formatter interface {
	newTimeStruct(value string, t time.Time) TimeStruct
	fromDateTimeString(value string) (time.Time, error)
	fromUnixTimestamp(value string) (time.Time, error)
}

var layout = time.RFC3339

// NewTimeStruct creates a new TimeStruct from string value
func NewTimeStruct(value string, formatter formatter) TimeStruct {
	t, e := formatter.fromUnixTimestamp(value)
	if e == nil {
		return formatter.newTimeStruct(value, t)
	}

	t2, e2 := formatter.fromDateTimeString(value)
	if e2 == nil {
		return formatter.newTimeStruct(value, t2)
	}

	return formatter.newTimeStruct("Current Time", time.Now())
}

// TimeStructFormatter TimeStructFormatter
type TimeStructFormatter struct{}

func (formatter TimeStructFormatter) newTimeStruct(v string, t time.Time) TimeStruct {
	return TimeStruct{
		Value:          v,
		Unixtime:       t.Unix(),
		UnixtimeMillis: t.UnixNano() / int64(time.Millisecond),
		DateTime:       t.UTC().Format(layout),
		LocalDateTime:  t.Local().Format(layout),
	}
}

func (formatter TimeStructFormatter) fromDateTimeString(value string) (time.Time, error) {
	if hhmmss.MatchString(value) {
		yymmdd := time.Now().Format("2006-01-02")
		return time.Parse(layout, yymmdd+"T"+value)
	}
	if mmdd.MatchString(value) {
		now := time.Now()
		year := now.Year()
		t, e := time.Parse(layout, fmt.Sprintf("%d-%s", year, value))
		if e == nil {
			return t, e
		}
		return time.Parse(layout, fmt.Sprintf("%d-%sT00:00:00Z", year, value[:5]))
	}
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
