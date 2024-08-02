package date

import (
	"fmt"
	"testing"
	"time"
)

func Test_presetFormatter_ToTime_today(t *testing.T) {
	sut := presetFormatter{
		input: "today",
	}

	actual, err := sut.ToTime()
	if err != nil {
		t.Error("error should not be thrown")
	}

	now := time.Now()
	expected := fmt.Sprintf("%d-%02d-%02dT00:00:00Z", now.Year(), now.Month(), now.Day())

	format := actual.UTC().Format(time.RFC3339)
	if format != expected {
		t.Errorf("assert failed. expect:%s, actual:%s", expected, format)
	}
}

func Test_presetFormatter_ToTimeStruct(t *testing.T) {
	sut := presetFormatter{
		input: "today",
	}

	v := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	actual := sut.ToTimeStruct(v)

	if actual.Value != "Today" {
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
