package date

import (
	"testing"
)

func Test_fromDateTimeString_success_mmdd(t *testing.T) {
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

func Test_fromDateTimeString_success_mmddhhmmss(t *testing.T) {
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
