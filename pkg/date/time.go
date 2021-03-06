package date

import "time"

func (formatter timeInputFormatter) ToTime() (time.Time, error) {
	yymmdd := time.Now().Format("2006-01-02")
	if formatter.precision == "second" {
		return time.Parse(layout, yymmdd+"T"+formatter.input+"Z")
	}
	return time.Parse(layout, yymmdd+"T"+formatter.input+":00Z")
}

func (formatter timeInputFormatter) ToTimeStruct(t time.Time) TimeStruct {
	return newTimeStruct(formatter.input, t)
}
