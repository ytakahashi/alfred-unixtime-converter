package date

import (
	"testing"
	"time"
)

func Test_currentTimeFormatter_ToTime(t *testing.T) {
	sut := currentTimeFormatter{}

	_, err := sut.ToTime()
	if err != nil {
		t.Error(err)
	}
}

func Test_currentTimeFormatter_ToTimeStruct(t *testing.T) {
	sut := currentTimeFormatter{}

	v := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	actual := sut.ToTimeStruct(v)

	if actual.Value != "Current Time" {
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
