package date

import (
	"reflect"
	"testing"
	"time"
)

func Test_NewInputFormatter_dateInputFormatter(t *testing.T) {
	input := "06-10"

	actual := NewInputFormatter(input)

	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(dateInputFormatter{})

	actualFormatter := actual.(dateInputFormatter)
	if actualType != expectedType {
		t.Errorf("Error. expect: %s actual: %s", expectedType, actualType)
	}
	if actualFormatter.input != input {
		t.Errorf("Error. expect: %s actual: %s", input, actualFormatter.input)
	}
}

func Test_NewInputFormatter_timeInputFormatter_second(t *testing.T) {
	input := "12:13:14"

	actual := NewInputFormatter(input)

	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(timeInputFormatter{})

	if actualType != expectedType {
		t.Errorf("Error. expect: %s actual: %s", expectedType, actualType)
	}

	actualFormatter := actual.(timeInputFormatter)
	if actualFormatter.input != input {
		t.Errorf("Error. expect: %s actual: %s", input, actualFormatter.input)
	}
	if actualFormatter.precision != "second" {
		t.Errorf("Error. actual: %s", actualFormatter.precision)
	}
}

func Test_NewInputFormatter_timeInputFormatter_minute(t *testing.T) {
	input := "12:13"

	actual := NewInputFormatter(input)

	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(timeInputFormatter{})

	if actualType != expectedType {
		t.Errorf("Error. expect: %s actual: %s", expectedType, actualType)
	}

	actualFormatter := actual.(timeInputFormatter)
	if actualFormatter.input != input {
		t.Errorf("Error. expect: %s actual: %s", input, actualFormatter.input)
	}
	if actualFormatter.precision != "minute" {
		t.Errorf("Error. actual: %s", actualFormatter.precision)
	}
}

func Test_NewInputFormatter_dateTimeInputFormatter(t *testing.T) {
	input := "2020-06-10T12:34:56Z"

	actual := NewInputFormatter(input)

	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(dateTimeInputFormatter{})

	actualFormatter := actual.(dateTimeInputFormatter)
	if actualType != expectedType {
		t.Errorf("Error. expect: %s actual: %s", expectedType, actualType)
	}
	if actualFormatter.input != input {
		t.Errorf("Error. expect: %s actual: %s", input, actualFormatter.input)
	}
}

func Test_NewInputFormatter_unixTimeInputFormatter(t *testing.T) {
	input := "1592724053"

	actual := NewInputFormatter(input)

	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(unixTimeInputFormatter{})

	actualFormatter := actual.(unixTimeInputFormatter)
	if actualType != expectedType {
		t.Errorf("Error. expect: %s actual: %s", expectedType, actualType)
	}
	if actualFormatter.input != input {
		t.Errorf("Error. expect: %s actual: %s", input, actualFormatter.input)
	}
}

func Test_NewInputFormatter_currentTimeFormatter(t *testing.T) {
	input := "a"

	actual := NewInputFormatter(input)

	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(currentTimeFormatter{})

	if actualType != expectedType {
		t.Errorf("Error. expect: %s actual: %s", expectedType, actualType)
	}
}

func Test_newTimeStruct(t *testing.T) {
	v := "test"
	tm := time.Unix(1293667200, 0)

	actual := newTimeStruct(v, tm)
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
