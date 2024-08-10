package date

import (
	"testing"
	"time"
)

func Test_unixTimeInputFormatter_ToTime_success(t *testing.T) {
	value := "1567299660"
	expected := "2019-09-01T01:01:00Z"

	sut := unixTimeInputFormatter{
		input: value,
		unit:  "s",
	}
	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thrown")
	}

	format := actual.UTC().Format(time.RFC3339)
	if format != expected {
		t.Errorf("assert failed. expect:%s, actual:%s", expected, format)
	}
}

func Test_unixTimeInputFormatter_ToTime_success_ms(t *testing.T) {
	value := "1567299660000"
	expected := "2019-09-01T01:01:00Z"

	sut := unixTimeInputFormatter{
		input: value,
		unit:  "ms",
	}
	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thrown")
	}

	format := actual.UTC().Format(time.RFC3339)
	if format != expected {
		t.Errorf("assert failed. expect:%s, actual:%s", expected, format)
	}
}

func Test_unixTimeInputFormatter_ToTimeStruct(t *testing.T) {
	input := "1257894000"
	sut := unixTimeInputFormatter{
		input: input,
		unit:  "s",
	}

	v := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	actual := sut.ToTimeStruct(v)

	if actual.Value != "1257894000 (s)" {
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
