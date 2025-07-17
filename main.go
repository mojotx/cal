package main

import (
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/mojotx/cal/pkg/calendar"
)

func main() {
	switch len(os.Args) {

	// No arguments provided
	case 1:
		now := time.Now()
		calendar.DumpMonth(now.Month(), now.Year())

	// One argument is a year
	case 2:
		year, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			color.Red("error parsing year: %s", err.Error())
			os.Exit(1)
		}
		calendar.DumpYear(int(year))

	// Two arguments: month and year
	case 3:
		month, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			color.Red("error parsing month: %s", err.Error())
			os.Exit(1)
		}
		year, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			color.Red("error parsing year: %s", err.Error())
			os.Exit(1)
		}
		calendar.DumpMonth(time.Month(month), int(year))

	default:
		color.Red("usage: %s [month] [year]", os.Args[0])

	}
}
