package alfred

import (
	"encoding/json"
	"strconv"

	"github.com/ytakahashi/alfred-unixtime-converter/date"
)

// Item is an alfred item (https://www.alfredapp.com/help/workflows/inputs/script-filter/json/)
type Item struct {
	UID      string `json:"uid"`
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

// Items ... array of items
type Items struct {
	Items []Item `json:"items"`
}

// ConvertToAlfredJSON converts a TimeStruct to JSON string
func ConvertToAlfredJSON(dt date.TimeStruct) string {
	datetime := dt.DateTime
	unixtime := strconv.FormatInt(dt.Unixtime, 10)
	unixtimems := strconv.FormatInt(dt.UnixtimeMillis, 10)

	item1 := Item{
		UID:      datetime,
		Title:    datetime,
		SubTitle: "DateTime string (ISO8601 format)",
		Arg:      datetime,
	}
	item2 := Item{
		UID:      unixtime,
		Title:    unixtime,
		SubTitle: "Unix Timestamp (s)",
		Arg:      unixtime,
	}
	item3 := Item{
		UID:      unixtimems,
		Title:    unixtimems,
		SubTitle: "Unix Timestamp (ms)",
		Arg:      unixtimems,
	}

	items := Items{Items: []Item{item1, item2, item3}}

	json, err := json.Marshal(items)
	if err != nil {
		return "{\"items\":[]}"
	}

	return string(json)
}
