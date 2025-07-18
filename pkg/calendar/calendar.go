package calendar

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

// NCenter centers a string in a buffer with a specified width.
func NCenter(width int, s string) *bytes.Buffer {
	const space = "\u0020"
	var b bytes.Buffer
	strLen := utf8.RuneCountInString(s)
	totalPad := width - strLen
	if totalPad < 1 {
		fmt.Fprint(&b, s)
		return &b
	}
	leftPad := totalPad / 2
	rightPad := totalPad - leftPad
	fmt.Fprintf(&b, "%s%s%s", strings.Repeat(space, leftPad), s, strings.Repeat(space, rightPad))
	return &b
}

// buildMonthCalendar generates a calendar for a specific month and year.
func buildMonthCalendar(month time.Month, year int) string {
	title := fmt.Sprintf("%s %d", month, year)
	b := NCenter(20, title)
	b.WriteRune('\n')
	b.WriteString("Su Mo Tu We Th Fr Sa\n")

	firstDayThisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	firstDayNextMonth := firstDayThisMonth.AddDate(0, 1, 0)
	lastDayThisMonth := firstDayNextMonth.AddDate(0, 0, -1)

	weekday := firstDayThisMonth.Weekday()
	Spacer(b, weekday)

	now := time.Now()
	todayYear, todayMonth, todayDay := now.Date()

	for day := firstDayThisMonth; day.Before(lastDayThisMonth) || day.Equal(lastDayThisMonth); day = day.AddDate(0, 0, 1) {
		dayStr := fmt.Sprintf("%2d", day.Day())
		if day.Year() == todayYear && day.Month() == todayMonth && day.Day() == todayDay {
			dayStr = color.New(color.BgWhite, color.FgBlack).Sprint(dayStr)
		}
		fmt.Fprintf(b, "%s ", dayStr)
		if day.Weekday() == time.Saturday {
			b.WriteRune('\n')
		}
	}
	b.WriteRune('\n')
	return b.String()
}

// DumpMonth prints the calendar for a specific month and year.
func DumpMonth(month time.Month, year int) {
	fmt.Print(buildMonthCalendar(month, year))
}

// DumpMonthToSlice returns the calendar for a specific month and year as a slice of strings.
func DumpMonthToSlice(month time.Month, year int) []string {
	calStr := buildMonthCalendar(month, year)
	var lineSlice []string
	scanner := bufio.NewScanner(strings.NewReader(calStr))
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\u0020\n\r\t")
		if line != "" {
			lineSlice = append(lineSlice, line)
		}
	}
	return lineSlice
}

// Spacer writes leading spaces to the buffer based on the weekday.
func Spacer(b *bytes.Buffer, weekday time.Weekday) {
	for n := int(weekday); n > 0; n-- {
		b.WriteString("\u0020\u0020\u0020")
	}
}

// dumpThreeMonths is a helper that prints three months in a row.
func dumpThreeMonths(year int, months ...time.Month) error {
	if len(months) != 3 {
		return errors.New("dumpThreeMonths requires exactly three months")
	}

	monthStrings := make(map[time.Month][]string)

	for i := months[0]; i <= months[2]; i++ {
		month := DumpMonthToSlice(i, year)
		monthStrings[i] = month
	}

	maxSliceLen := GetMaxSliceLen(monthStrings[months[0]], monthStrings[months[1]], monthStrings[months[2]])

	for i := range maxSliceLen {
		for _, month := range months {
			// Check for a line for this month
			var subString string
			if i <= len(monthStrings[month])-1 {
				subString = monthStrings[month][i]
			} else {
				subString = strings.Repeat(" ", 20) // Empty space for alignment
			}
			fmt.Printf("%-20s    ", subString)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
	return nil
}

// GetMaxSliceLen returns the maximum length of the provided slices.
func GetMaxSliceLen(slices ...[]string) int {
	max := math.MinInt
	for _, slice := range slices {
		if len(slice) > max {
			max = len(slice)
		}
	}
	return max
}

// DumpYear prints the calendar for an entire year.
func DumpYear(year int) {
	_ = dumpThreeMonths(year, time.January, time.February, time.March)
	_ = dumpThreeMonths(year, time.April, time.May, time.June)
	_ = dumpThreeMonths(year, time.July, time.August, time.September)
	_ = dumpThreeMonths(year, time.October, time.November, time.December)
}

// Helper function to strip ANSI color codes for testing
func stripAnsiCodes(s string) string {
	ansi := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return ansi.ReplaceAllString(s, "")
}
