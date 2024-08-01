package date

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

// TimeStruct is a struct
type TimeStruct struct {
	Value          string
	Unixtime       int64
	UnixtimeMillis int64
	DateTime       string
	LocalDateTime  string
}

type presetString string

const (
	NOW   = presetString("now")
	TODAY = presetString("today")
)

var preset = []string{string(NOW), string(TODAY)}
var hhmmss = regexp.MustCompile(`^[0-2]?[0-9]:[0-5][0-9]:[0-5][0-9]`)
var hhmm = regexp.MustCompile(`^[0-2]?[0-9]:[0-5][0-9]$`)
var mmdd = regexp.MustCompile(`^[0-1][0-9]\-[0-3][0-9]`)
var year = regexp.MustCompile(`^[1-2][0-9]{3}\-`)
var layout = time.RFC3339

// InputFormatter InputFormatter
type InputFormatter interface {
	ToTime() (time.Time, error)
	ToTimeStruct(t time.Time) TimeStruct
}

type presetFormatter struct {
	input string
}
type dateTimeInputFormatter struct {
	input string
}
type unixTimeInputFormatter struct {
	input string
}
type dateInputFormatter struct {
	input string
}
type timeInputFormatter struct {
	input     string
	precision string
}
type currentTimeFormatter struct{}

// NewInputFormatter creates a new input formatter instance
func NewInputFormatter(input string) InputFormatter {
	if slices.Contains(preset, strings.ToLower(input)) {
		return presetFormatter{
			input: input,
		}
	}
	if mmdd.MatchString(input) {
		return dateInputFormatter{
			input: input,
		}
	}
	if hhmmss.MatchString(input) {
		return timeInputFormatter{
			input:     hhmmss.FindString(input),
			precision: "second",
		}
	}
	if hhmm.MatchString(input) {
		return timeInputFormatter{
			input:     input,
			precision: "minute",
		}
	}
	if year.MatchString(input) {
		return dateTimeInputFormatter{
			input: input,
		}
	}
	if _, e := strconv.ParseInt(input, 10, 64); e == nil {
		return unixTimeInputFormatter{
			input: input,
		}
	}
	return currentTimeFormatter{}
}

func newTimeStruct(v string, t time.Time) TimeStruct {
	return TimeStruct{
		Value:          v,
		Unixtime:       t.Unix(),
		UnixtimeMillis: t.UnixNano() / int64(time.Millisecond),
		DateTime:       t.UTC().Format(layout),
		LocalDateTime:  t.Local().Format(layout),
	}
}
