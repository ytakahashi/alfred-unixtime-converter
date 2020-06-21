package date

import (
	"testing"
	"time"
)

func Test_dateTimeInputFormatter_ToTime_error(t *testing.T) {
	value := "abc"

	sut := dateTimeInputFormatter{
		input: value,
	}
	_, err := sut.ToTime()
	if err == nil {
		t.Error("error should be thorown")
	}
}

func Test_dateTimeInputFormatter_ToTime_success(t *testing.T) {
	value := "2019-01-02T03:04:00Z"
	expected := "2019-01-02T03:04:00Z"

	sut := dateTimeInputFormatter{
		input: value,
	}
	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thorown")
	}
	format := actual.UTC().Format(time.RFC3339)
	if format != expected {
		t.Errorf("assert failed. expect:%s, actual:%s", expected, format)
	}
}

func Test_dateTimeInputFormatter_ToTime_success_local(t *testing.T) {
	value := "2019-01-02T20:04:00+09:00"
	expected := "2019-01-02T11:04:00Z"

	sut := dateTimeInputFormatter{
		input: value,
	}
	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thorown")
	}
	format := actual.UTC().Format(time.RFC3339)
	if format != expected {
		t.Errorf("assert failed. expect:%s, actual:%s", expected, format)
	}
}

func Test_dateTimeInputFormatter_ToTimeStruct(t *testing.T) {
	input := "2019-01-02T03:04:00Z"
	sut := dateTimeInputFormatter{
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
