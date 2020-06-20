package date

import "testing"

func Test_ToTime_hhmmss(t *testing.T) {
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

func Test_ToTime_hhmmss_local(t *testing.T) {
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

func Test_ToTime_success_hhmm(t *testing.T) {
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
