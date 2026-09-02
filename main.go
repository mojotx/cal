package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/mojotx/cal/pkg/calendar"
)

func parseIntArg(s string) (int, error) {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	if value < math.MinInt || value > math.MaxInt {
		return 0, fmt.Errorf("value %q overflows int", s)
	}
	return int(value), nil
}

func main() {
	switch len(os.Args) {

	// No arguments provided
	case 1:
		now := time.Now()
		calendar.DumpMonth(now.Month(), now.Year())

	// One argument is a year
	case 2:
		year, err := parseIntArg(os.Args[1])
		if err != nil {
			color.Red("error parsing year: %s", err.Error())
			os.Exit(1)
		}
		if err := calendar.ValidateYear(year); err != nil {
			color.Red("error: %s", err.Error())
			os.Exit(1)
		}
		if err := calendar.DumpYear(year); err != nil {
			color.Red("error: %s", err.Error())
			os.Exit(1)
		}

	// Two arguments: month and year
	case 3:
		month, err := parseIntArg(os.Args[1])
		if err != nil {
			color.Red("error parsing month: %s", err.Error())
			os.Exit(1)
		}
		year, err := parseIntArg(os.Args[2])
		if err != nil {
			color.Red("error parsing year: %s", err.Error())
			os.Exit(1)
		}
		if err := calendar.ValidateMonth(month); err != nil {
			color.Red("error: %s", err.Error())
			os.Exit(1)
		}
		if err := calendar.ValidateYear(year); err != nil {
			color.Red("error: %s", err.Error())
			os.Exit(1)
		}
		calendar.DumpMonth(time.Month(month), year)

	default:
		color.Red("usage: %s [month] [year]", os.Args[0])

	}
}
