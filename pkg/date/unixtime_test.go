package date

import (
	"testing"
	"time"
)

func Test_fromUnixTimestamp_success(t *testing.T) {
	value := "1567299660"
	expected := "2019-09-01T01:01:00Z"

	sut := unixTimeInputFormatter{
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
