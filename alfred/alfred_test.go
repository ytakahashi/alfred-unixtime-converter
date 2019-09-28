package alfred

import (
	"testing"

	"github.com/ytakahashi/alfred-unixtime-converter/date"
)

func Test_Convert(t *testing.T) {
	value := date.TimeStruct{
		DateTime:       "2020-01-01T00:00:00Z",
		Unixtime:       1577836800,
		UnixtimeMillis: 1577836800000,
	}

	expected := "{\"items\":[{\"uid\":\"2020-01-01T00:00:00Z\",\"title\":\"2020-01-01T00:00:00Z\",\"subtitle\":\"DateTime string (ISO8601 format)\",\"arg\":\"2020-01-01T00:00:00Z\"},{\"uid\":\"1577836800\",\"title\":\"1577836800\",\"subtitle\":\"Unix Timestamp (s)\",\"arg\":\"1577836800\"},{\"uid\":\"1577836800000\",\"title\":\"1577836800000\",\"subtitle\":\"Unix Timestamp (ms)\",\"arg\":\"1577836800000\"}]}"
	actual := ConvertToAlfredJSON(value)
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
	}
}
