package date

import (
	"testing"
	"time"
)

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
