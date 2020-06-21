package date

import (
	"testing"
	"time"
)

func Test_dateInputFormatter_ToTime_success_mmdd(t *testing.T) {
	value := "04-01"

	sut := dateInputFormatter{
		input: value,
	}
	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thorown")
	}

	year, month, day := actual.Date()
	if year < 2020 {
		t.Errorf("unexpected hour: %d", year)
	}
	if month != 4 {
		t.Errorf("unexpected month: %d", month)
	}
	if day != 1 {
		t.Errorf("unexpected day: %d", day)
	}
}

func Test_dateInputFormatter_ToTime_success_mmddhhmmss(t *testing.T) {
	value := "04-01T01:02:03Z"

	sut := dateInputFormatter{
		input: value,
	}
	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thorown")
	}

	year, month, day := actual.Date()
	if year < 2020 {
		t.Errorf("unexpected year: %d", year)
	}
	if month != 4 {
		t.Errorf("unexpected month: %d", month)
	}
	if day != 1 {
		t.Errorf("unexpected day: %d", day)
	}

	hour, min, sec := actual.Clock()
	if hour != 1 {
		t.Errorf("unexpected hour: %d", hour)
	}
	if min != 2 {
		t.Errorf("unexpected min: %d", min)
	}
	if sec != 3 {
		t.Errorf("unexpected sec: %d", sec)
	}
}

func Test_dateInputFormatter_ToTimeStruct(t *testing.T) {
	input := "11-12"
	sut := dateInputFormatter{
		input: input,
	}

	v := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	actual := sut.ToTimeStruct(v)

	if actual.Value != input {
		t.Errorf("unexpected Value (%s)", actual.Value)
	}
	if actual.DateTime != "2009-11-10T23:00:00Z" {
		t.Errorf("unexpected DateTimelue (%s)", actual.DateTime)
	}
	if actual.LocalDateTime == "" {
		t.Error("unexpected LocalDateTime")
	}
	if actual.Unixtime != 1257894000 {
		t.Errorf("unexpected Unixtime (%d)", actual.Unixtime)
	}
	if actual.UnixtimeMillis != 1257894000000 {
		t.Errorf("unexpected Unixtime (%d)", actual.UnixtimeMillis)
	}
}
