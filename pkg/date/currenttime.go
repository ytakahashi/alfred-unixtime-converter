package date

import "time"

func (formatter currentTimeFormatter) ToTime() (time.Time, error) {
	return time.Now(), nil
}

func (formatter currentTimeFormatter) ToTimeStruct(t time.Time) TimeStruct {
	return newTimeStruct("Current Time", t)
}
