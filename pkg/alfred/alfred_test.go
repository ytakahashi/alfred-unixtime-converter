package alfred

import (
	"testing"

	"github.com/ytakahashi/alfred-unixtime-converter/pkg/date"
)

func Test_Convert(t *testing.T) {
	value := date.TimeStruct{
		Value:          "test",
		DateTime:       "2020-01-01T00:00:00",
		LocalDateTime:  "2020-01-01T09:00:00+09:00",
		Unixtime:       1577836800,
		UnixtimeMillis: 1577836800000,
	}

	expected := `{"items":[{"uid":"2020-01-01T09:00:00+09:00","title":"2020-01-01T09:00:00+09:00","subtitle":"test: DateTime string (ISO8601, Local)","arg":"2020-01-01T09:00:00+09:00"},{"uid":"2020-01-01T00:00:00","title":"2020-01-01T00:00:00","subtitle":"test: DateTime string (ISO8601, UTC)","arg":"2020-01-01T00:00:00"},{"uid":"1577836800","title":"1577836800","subtitle":"test: Unix Timestamp (s)","arg":"1577836800"},{"uid":"1577836800000","title":"1577836800000","subtitle":"test: Unix Timestamp (ms)","arg":"1577836800000"}]}`
	actual := ConvertToAlfredJSON(value)
	if actual != expected {
		t.Errorf("assert failed. expect:%v actual:%v", expected, actual)
	}
}
