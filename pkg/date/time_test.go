package date

import (
	"testing"
	"time"
)

func Test_timeInputFormatter_ToTime_hhmmss(t *testing.T) {
	value := "01:01:00Z"

	sut := timeInputFormatter{
		input:     value,
		precision: "second",
	}
	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thorown")
	}

	hour, min, sec := actual.Clock()
	if hour != 1 {
		t.Errorf("unexpected hour: %d", hour)
	}
	if min != 1 {
		t.Errorf("unexpected min: %d", min)
	}
	if sec != 0 {
		t.Errorf("unexpected sec: %d", sec)
	}
}

func Test_timeInputFormatter_ToTime_hhmmss_local(t *testing.T) {
	value := "10:01:00+09:00"

	sut := timeInputFormatter{
		input:     value,
		precision: "second",
	}
	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thorown")
	}

	hour, min, sec := actual.UTC().Clock()
	if hour != 1 {
		t.Errorf("unexpected hour: %d", hour)
	}
	if min != 1 {
		t.Errorf("unexpected min: %d", min)
	}
	if sec != 0 {
		t.Errorf("unexpected sec: %d", sec)
	}
}

func Test_timeInputFormatter_ToTime_success_hhmm(t *testing.T) {
	value := "02:02"

	sut := timeInputFormatter{
		input:     value,
		precision: "minute",
	}
	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thorown")
	}

	hour, min, sec := actual.Clock()
	if hour != 2 {
		t.Errorf("unexpected hour: %d", hour)
	}
	if min != 2 {
		t.Errorf("unexpected min: %d", min)
	}
	if sec != 0 {
		t.Errorf("unexpected sec: %d", sec)
	}
}

func Test_timeInputFormatter_ToTimeStruct(t *testing.T) {
	input := "10:45"
	sut := timeInputFormatter{
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
