package date

import "time"

func (formatter dateTimeInputFormatter) ToTime() (time.Time, error) {
	return time.Parse(layout, formatter.input)
}

func (formatter dateTimeInputFormatter) ToTimeStruct(t time.Time) TimeStruct {
	return newTimeStruct(formatter.input, t)
}
