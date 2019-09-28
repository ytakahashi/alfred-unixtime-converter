package main

import (
	"flag"
	"fmt"

	"github.com/ytakahashi/alfred-unixtime-converter/alfred"
	"github.com/ytakahashi/alfred-unixtime-converter/date"
)

type timeFormatter interface {
	Format(value string) date.TimeStruct
}

type timeConverter interface {
	Convert(value date.TimeStruct) string
}

func main() {
	flag.Parse()
	val := flag.Arg(0)

	formatter := date.TimeStructFormatter{}
	time := date.NewTimeStruct(val, formatter)

	st := alfred.ConvertToAlfredJSON(time)

	fmt.Println(st)
}
