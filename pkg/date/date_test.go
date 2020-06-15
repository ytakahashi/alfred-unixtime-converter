package date

import (
	"errors"
	"testing"
	"time"
)

var expected = TimeStruct{
	Unixtime:       1293667200,
	UnixtimeMillis: 1293667200000,
	DateTime:       "2010-12-30T00:00:00Z",
}

type mockedTimeStructFormatter1 struct{}

func (formatter mockedTimeStructFormatter1) newTimeStruct(v string, t time.Time) TimeStruct {
	return expected
}

func (formatter mockedTimeStructFormatter1) fromDateTimeString(value string) (time.Time, error) {
	if value == "2010-12-30T00:00:00Z" {
		return time.Now(), nil
	}
	return time.Now(), errors.New("error")
}

func (formatter mockedTimeStructFormatter1) fromUnixTimestamp(value string) (time.Time, error) {
	if value == "1293667200" {
		return time.Now(), nil
	}
	return time.Now(), errors.New("error")
}

func Test_NewTimeStruct_fromDateTime(t *testing.T) {
	value := "2010-12-30T00:00:00Z"
	formatter := mockedTimeStructFormatter1{}

	actual := NewTimeStruct(value, formatter)
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
	}
}

func Test_NewTimeStruct_fromTimestamp(t *testing.T) {
	value := "1293667200"
	formatter := mockedTimeStructFormatter1{}

	actual := NewTimeStruct(value, formatter)
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
	}
}

func Test_NewTimeStruct_now(t *testing.T) {
	value := ""
	formatter := mockedTimeStructFormatter1{}

	actual := NewTimeStruct(value, formatter)
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
	}
}
