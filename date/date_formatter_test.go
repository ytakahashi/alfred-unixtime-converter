package date

import (
	"testing"
	"time"
)

func Test_newTimeStruct(t *testing.T) {
	v := time.Unix(1293667200, 0)

	sut := TimeStructFormatter{}
	actual := sut.newTimeStruct(v)
	expected := TimeStruct{
		Unixtime:       1293667200,
		UnixtimeMillis: 1293667200000,
		DateTime:       "2010-12-30T00:00:00Z",
	}
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
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
