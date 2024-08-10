package main

import (
	"flag"
	"fmt"

	"github.com/ytakahashi/alfred-unixtime-converter/pkg/alfred"
	"github.com/ytakahashi/alfred-unixtime-converter/pkg/date"
)

func main() {
	flag.Parse()
	input := flag.Arg(0)
	option := flag.Arg(1)

	formatter := date.NewInputFormatter(input, option)
	execute(formatter)
}

func execute(formatter date.InputFormatter) {
	t, err := formatter.ToTime()
	if err != nil {
		fmt.Println(err)
	}
	time := formatter.ToTimeStruct(t)
	st := alfred.ConvertToAlfredJSON(time)

	fmt.Println(st)
}
