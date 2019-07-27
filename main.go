package main

import (
	"flag"
	"fmt"

	"github.com/ytakahashi/alfred-unixtime-converter/alfred"
	"github.com/ytakahashi/alfred-unixtime-converter/date"
)

func main() {
	flag.Parse()
	val := flag.Arg(0)

	time := date.NewTimeStruct(val)

	st := alfred.ConvertToAlfredJSON(time)

	fmt.Println(st)
}
