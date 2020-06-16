package date

import (
	"testing"
	"time"
)

func Test_fromDateTimeString_error(t *testing.T) {
	value := "abc"

	sut := dateTimeInputFormatter{
		input: value,
	}
	_, err := sut.ToTime()
	if err == nil {
		t.Error("error should be thorown")
	}
}

func Test_fromDateTimeString_success(t *testing.T) {
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

func Test_fromDateTimeString_success_local(t *testing.T) {
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
