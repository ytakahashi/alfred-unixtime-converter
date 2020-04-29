package date

import (
	"testing"
	"time"
)

func Test_newTimeStruct(t *testing.T) {
	v := "test"
	tm := time.Unix(1293667200, 0)

	sut := TimeStructFormatter{}
	actual := sut.newTimeStruct(v, tm)
	if actual.Value != v {
		t.Errorf("Unexpected Value. expect:%s actual:%s", v, actual.Value)
	}
	if actual.Unixtime != 1293667200 {
		t.Errorf("Unexpected Unixtime. expect:%d actual:%d", 1293667200, actual.Unixtime)
	}
	if actual.UnixtimeMillis != 1293667200000 {
		t.Errorf("Unexpected UnixtimeMillis. expect:%d actual:%d", 1293667200000, actual.UnixtimeMillis)
	}
	if actual.DateTime != "2010-12-30T00:00:00Z" {
		t.Errorf("Unexpected DateTime. expect:%s actual:%s", "2010-12-30T00:00:00Z", actual.DateTime)
	}
	if actual.LocalDateTime == "" {
		t.Error("Unexpected LocalDateTime.")
	}
}

func Test_fromDateTimeString_error(t *testing.T) {
	value := "abc"

	sut := TimeStructFormatter{}
	_, err := sut.fromDateTimeString(value)
	if err == nil {
		t.Error("error should be thorown")
	}
}

func Test_fromDateTimeString_success(t *testing.T) {
	value := "2019-01-02T03:04:00Z"
	expected := "2019-01-02T03:04:00Z"

	sut := TimeStructFormatter{}
	actual, err := sut.fromDateTimeString(value)
	if err != nil {
		t.Error("error should not be thorown")
	}
	format := actual.UTC().Format(time.RFC3339)
	if format != expected {
		t.Errorf("assert failed. expect:%s, actual:%s", expected, format)
	}
}

func Test_fromDateTimeString_success_hhmmss(t *testing.T) {
	value := "01:01:00Z"

	sut := TimeStructFormatter{}
	actual, err := sut.fromDateTimeString(value)
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

func Test_fromDateTimeString_success_mmdd(t *testing.T) {
	value := "04-01"

	sut := TimeStructFormatter{}
	actual, err := sut.fromDateTimeString(value)
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

func Test_fromDateTimeString_success_mmddhhmmss(t *testing.T) {
	value := "04-01T01:02:03Z"

	sut := TimeStructFormatter{}
	actual, err := sut.fromDateTimeString(value)
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

func Test_fromUnixTimestamp_error(t *testing.T) {
	value := "abc"

	sut := TimeStructFormatter{}
	_, err := sut.fromUnixTimestamp(value)
	if err == nil {
		t.Error("error should be thorown")
	}
}

func Test_fromUnixTimestamp_success(t *testing.T) {
	value := "1567299660"
	expected := "2019-09-01T01:01:00Z"

	sut := TimeStructFormatter{}
	actual, err := sut.fromUnixTimestamp(value)
	if err != nil {
		t.Error("error should not be thorown")
	}

	format := actual.UTC().Format(time.RFC3339)
	if format != expected {
		t.Errorf("assert failed. expect:%s, actual:%s", expected, format)
	}
}
