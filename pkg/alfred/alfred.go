package alfred

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ytakahashi/alfred-unixtime-converter/pkg/date"
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

var emptyItems = "{\"items\":[]}"

// ConvertToAlfredJSON converts a TimeStruct to JSON string
func ConvertToAlfredJSON(dt date.TimeStruct) string {
	value := dt.Value
	localDatetime := dt.LocalDateTime
	datetime := dt.DateTime
	unixtime := strconv.FormatInt(dt.Unixtime, 10)
	unixtimems := strconv.FormatInt(dt.UnixtimeMillis, 10)

	item0 := Item{
		UID:      localDatetime,
		Title:    localDatetime,
		SubTitle: fmt.Sprintf("%s: DateTime string (ISO8601, Local)", value),
		Arg:      localDatetime,
	}
	item1 := Item{
		UID:      datetime,
		Title:    datetime,
		SubTitle: fmt.Sprintf("%s: DateTime string (ISO8601, UTC)", value),
		Arg:      datetime,
	}
	item2 := Item{
		UID:      unixtime,
		Title:    unixtime,
		SubTitle: fmt.Sprintf("%s: Unix Timestamp (s)", value),
		Arg:      unixtime,
	}
	item3 := Item{
		UID:      unixtimems,
		Title:    unixtimems,
		SubTitle: fmt.Sprintf("%s: Unix Timestamp (ms)", value),
		Arg:      unixtimems,
	}

	items := Items{Items: []Item{item0, item1, item2, item3}}

	json, err := json.Marshal(items)
	if err != nil {
		return emptyItems
	}

	return string(json)
}
