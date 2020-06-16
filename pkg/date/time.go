package date

import "time"

func (formatter timeInputFormatter) ToTime() (time.Time, error) {
	yymmdd := time.Now().Format("2006-01-02")
	return time.Parse(layout, yymmdd+"T"+formatter.input)
}

func (formatter timeInputFormatter) ToTimeStruct(t time.Time) TimeStruct {
	return newTimeStruct(formatter.input, t)
}
