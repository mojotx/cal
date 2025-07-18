package calendar

import (
	"bytes"
	"math"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNCenter(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		input    string
		expected string
	}{
		{
			name:     "normal centering",
			width:    10,
			input:    "test",
			expected: "   test   ",
		},
		{
			name:     "odd string length",
			width:    10,
			input:    "hello",
			expected: "  hello   ",
		},
		{
			name:     "width equal to string length",
			width:    4,
			input:    "test",
			expected: "test",
		},
		{
			name:     "width less than string length",
			width:    2,
			input:    "test",
			expected: "test",
		},
		{
			name:     "empty string",
			width:    5,
			input:    "",
			expected: "     ",
		},
		{
			name:     "odd width even string",
			width:    7,
			input:    "ab",
			expected: "  ab   ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NCenter(tt.width, tt.input).String()
			assert.Equal(t, tt.expected, result, "NCenter should return the expected centered string")
		})
	}
}
func TestBuildMonthCalendar(t *testing.T) {
	tests := []struct {
		name     string
		month    time.Month
		year     int
		expected string
	}{
		{
			name:  "July 2025",
			month: time.July,
			year:  2025,
			expected: "     July 2025      \n" +
				"Su Mo Tu We Th Fr Sa\n" +
				"       1  2  3  4  5 \n" +
				" 6  7  8  9 10 11 12 \n" +
				"13 14 15 16 17 18 19 \n" +
				"20 21 22 23 24 25 26 \n" +
				"27 28 29 30 31 \n",
		},
		{
			name:  "February 2024 (leap year)",
			month: time.February,
			year:  2024,
			expected: "   February 2024    \n" +
				"Su Mo Tu We Th Fr Sa\n" +
				"             1  2  3 \n" +
				" 4  5  6  7  8  9 10 \n" +
				"11 12 13 14 15 16 17 \n" +
				"18 19 20 21 22 23 24 \n" +
				"25 26 27 28 29 \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildMonthCalendar(tt.month, tt.year)

			// Remove any ANSI color codes for comparison
			result = stripAnsiCodes(result)

			assert.Equal(t, tt.expected, result, "Calendar output should match expected format")
		})
	}
}
func TestStripAnsiCodes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "string with no ANSI codes",
			input:    "plain text",
			expected: "plain text",
		},
		{
			name:     "string with ANSI color code",
			input:    "\x1b[31mred text\x1b[0m",
			expected: "red text",
		},
		{
			name:     "string with multiple ANSI codes",
			input:    "\x1b[1;31mbold red\x1b[0m \x1b[32mgreen\x1b[0m",
			expected: "bold red green",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "just ANSI codes",
			input:    "\x1b[31m\x1b[1m\x1b[0m",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stripAnsiCodes(tt.input)
			assert.Equal(t, tt.expected, result, "stripAnsiCodes should remove all ANSI color codes")
		})
	}
}
func TestDumpMonth(t *testing.T) {
	tests := []struct {
		name     string
		month    time.Month
		year     int
		expected string
	}{
		{
			name:  "Dump July 2025",
			month: time.July,
			year:  2025,
			expected: "     July 2025      \n" +
				"Su Mo Tu We Th Fr Sa\n" +
				"       1  2  3  4  5 \n" +
				" 6  7  8  9 10 11 12 \n" +
				"13 14 15 16 17 18 19 \n" +
				"20 21 22 23 24 25 26 \n" +
				"27 28 29 30 31 \n",
		},
		{
			name:  "Dump February 2024 (leap year)",
			month: time.February,
			year:  2024,
			expected: "   February 2024    \n" +
				"Su Mo Tu We Th Fr Sa\n" +
				"             1  2  3 \n" +
				" 4  5  6  7  8  9 10 \n" +
				"11 12 13 14 15 16 17 \n" +
				"18 19 20 21 22 23 24 \n" +
				"25 26 27 28 29 \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture stdout
			var buf bytes.Buffer
			stdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			DumpMonth(tt.month, tt.year)

			w.Close()
			os.Stdout = stdout
			if _, err := buf.ReadFrom(r); err != nil {
				t.Fatalf("Failed to read from pipe: %v", err)
			}
			output := stripAnsiCodes(buf.String())

			assert.Equal(t, tt.expected, output, "DumpMonth output should match expected calendar")
		})
	}
}
func TestDumpMonthToSlice(t *testing.T) {
	tests := []struct {
		name     string
		month    time.Month
		year     int
		expected []string
	}{
		{
			name:  "July 2025",
			month: time.July,
			year:  2025,
			expected: []string{
				"     July 2025",
				"Su Mo Tu We Th Fr Sa",
				"       1  2  3  4  5",
				" 6  7  8  9 10 11 12",
				"13 14 15 16 17 18 19",
				"20 21 22 23 24 25 26",
				"27 28 29 30 31",
			},
		},
		{
			name:  "February 2024 (leap year)",
			month: time.February,
			year:  2024,
			expected: []string{
				"   February 2024",
				"Su Mo Tu We Th Fr Sa",
				"             1  2  3",
				" 4  5  6  7  8  9 10",
				"11 12 13 14 15 16 17",
				"18 19 20 21 22 23 24",
				"25 26 27 28 29",
			},
		},
		{
			name:  "January 2023",
			month: time.January,
			year:  2023,
			expected: []string{
				"    January 2023",
				"Su Mo Tu We Th Fr Sa",
				" 1  2  3  4  5  6  7",
				" 8  9 10 11 12 13 14",
				"15 16 17 18 19 20 21",
				"22 23 24 25 26 27 28",
				"29 30 31",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DumpMonthToSlice(tt.month, tt.year)
			// Remove any ANSI color codes for comparison
			for i := range result {
				result[i] = stripAnsiCodes(result[i])
			}
			assert.Equal(t, tt.expected, result, "DumpMonthToSlice output should match expected slice")
		})
	}
}
func TestSpacer(t *testing.T) {
	tests := []struct {
		name     string
		weekday  time.Weekday
		expected string
	}{
		{
			name:     "Sunday (no spaces)",
			weekday:  time.Sunday,
			expected: "",
		},
		{
			name:     "Monday (3 spaces)",
			weekday:  time.Monday,
			expected: "   ",
		},
		{
			name:     "Wednesday (9 spaces)",
			weekday:  time.Wednesday,
			expected: "         ",
		},
		{
			name:     "Saturday (18 spaces)",
			weekday:  time.Saturday,
			expected: "                  ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			Spacer(&buf, tt.weekday)
			assert.Equal(t, tt.expected, buf.String(), "Spacer should write correct number of spaces")
		})
	}
}
func TestGetMaxSliceLen(t *testing.T) {
	tests := []struct {
		name     string
		slices   [][]string
		expected int
	}{
		{
			name:     "all slices empty",
			slices:   [][]string{{}, {}, {}},
			expected: 0,
		},
		{
			name:     "one slice longer",
			slices:   [][]string{{"a"}, {"b", "c"}, {"d", "e", "f"}},
			expected: 3,
		},
		{
			name:     "all slices same length",
			slices:   [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}},
			expected: 2,
		},
		{
			name:     "single slice",
			slices:   [][]string{{"x", "y", "z"}},
			expected: 3,
		},
		{
			name:     "no slices",
			slices:   [][]string{},
			expected: math.MinInt,
		},
		{
			name:     "mix of empty and non-empty",
			slices:   [][]string{{}, {"a"}, {}, {"b", "c"}},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetMaxSliceLen(tt.slices...)
			assert.Equal(t, tt.expected, result, "GetMaxSliceLen should return the correct maximum length")
		})
	}
}
func TestDumpThreeMonths(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		months   []time.Month
		wantErr  bool
		expected string
	}{
		{
			name:    "wrong number of months",
			year:    2023,
			months:  []time.Month{time.January, time.February},
			wantErr: true,
		},
		{
			name:   "first quarter",
			year:   2023,
			months: []time.Month{time.January, time.February, time.March},
			expected: "    January 2023           February 2023             March 2023         \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				" 1  2  3  4  5  6  7              1  2  3  4              1  2  3  4    \n" +
				" 8  9 10 11 12 13 14     5  6  7  8  9 10 11     5  6  7  8  9 10 11    \n" +
				"15 16 17 18 19 20 21    12 13 14 15 16 17 18    12 13 14 15 16 17 18    \n" +
				"22 23 24 25 26 27 28    19 20 21 22 23 24 25    19 20 21 22 23 24 25    \n" +
				"29 30 31                26 27 28                26 27 28 29 30 31       \n\n",
		},
		{
			name:   "second quarter",
			year:   2023,
			months: []time.Month{time.April, time.May, time.June},
			expected: "     April 2023               May 2023               June 2023          \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				"                   1        1  2  3  4  5  6                 1  2  3    \n" +
				" 2  3  4  5  6  7  8     7  8  9 10 11 12 13     4  5  6  7  8  9 10    \n" +
				" 9 10 11 12 13 14 15    14 15 16 17 18 19 20    11 12 13 14 15 16 17    \n" +
				"16 17 18 19 20 21 22    21 22 23 24 25 26 27    18 19 20 21 22 23 24    \n" +
				"23 24 25 26 27 28 29    28 29 30 31             25 26 27 28 29 30       \n" +
				"30                                                                      \n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			stdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := dumpThreeMonths(tt.year, tt.months...)

			w.Close()
			os.Stdout = stdout
			if _, readErr := buf.ReadFrom(r); readErr != nil {
				t.Fatalf("Failed to read from pipe: %v", readErr)
			}

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			output := stripAnsiCodes(buf.String())
			assert.Equal(t, tt.expected, output)
		})
	}
}
func TestDumpYear(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		expected string
	}{
		{
			name: "test year 2023",
			year: 2023,
			expected: "    January 2023           February 2023             March 2023         \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				" 1  2  3  4  5  6  7              1  2  3  4              1  2  3  4    \n" +
				" 8  9 10 11 12 13 14     5  6  7  8  9 10 11     5  6  7  8  9 10 11    \n" +
				"15 16 17 18 19 20 21    12 13 14 15 16 17 18    12 13 14 15 16 17 18    \n" +
				"22 23 24 25 26 27 28    19 20 21 22 23 24 25    19 20 21 22 23 24 25    \n" +
				"29 30 31                26 27 28                26 27 28 29 30 31       \n\n" +
				"     April 2023               May 2023               June 2023          \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				"                   1        1  2  3  4  5  6                 1  2  3    \n" +
				" 2  3  4  5  6  7  8     7  8  9 10 11 12 13     4  5  6  7  8  9 10    \n" +
				" 9 10 11 12 13 14 15    14 15 16 17 18 19 20    11 12 13 14 15 16 17    \n" +
				"16 17 18 19 20 21 22    21 22 23 24 25 26 27    18 19 20 21 22 23 24    \n" +
				"23 24 25 26 27 28 29    28 29 30 31             25 26 27 28 29 30       \n" +
				"30                                                                      \n\n" +
				"     July 2023              August 2023            September 2023       \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				"                   1           1  2  3  4  5                    1  2    \n" +
				" 2  3  4  5  6  7  8     6  7  8  9 10 11 12     3  4  5  6  7  8  9    \n" +
				" 9 10 11 12 13 14 15    13 14 15 16 17 18 19    10 11 12 13 14 15 16    \n" +
				"16 17 18 19 20 21 22    20 21 22 23 24 25 26    17 18 19 20 21 22 23    \n" +
				"23 24 25 26 27 28 29    27 28 29 30 31          24 25 26 27 28 29 30    \n" +
				"30 31                                                                   \n\n" +
				"    October 2023           November 2023           December 2023        \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				" 1  2  3  4  5  6  7              1  2  3  4                    1  2    \n" +
				" 8  9 10 11 12 13 14     5  6  7  8  9 10 11     3  4  5  6  7  8  9    \n" +
				"15 16 17 18 19 20 21    12 13 14 15 16 17 18    10 11 12 13 14 15 16    \n" +
				"22 23 24 25 26 27 28    19 20 21 22 23 24 25    17 18 19 20 21 22 23    \n" +
				"29 30 31                26 27 28 29 30          24 25 26 27 28 29 30    \n" +
				"                                                31                      \n\n",
		},
		{
			name: "test year 2024 (leap Year)",
			year: 2024,
			expected: "    January 2024           February 2024             March 2024         \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				"    1  2  3  4  5  6                 1  2  3                    1  2    \n" +
				" 7  8  9 10 11 12 13     4  5  6  7  8  9 10     3  4  5  6  7  8  9    \n" +
				"14 15 16 17 18 19 20    11 12 13 14 15 16 17    10 11 12 13 14 15 16    \n" +
				"21 22 23 24 25 26 27    18 19 20 21 22 23 24    17 18 19 20 21 22 23    \n" +
				"28 29 30 31             25 26 27 28 29          24 25 26 27 28 29 30    \n" +
				"                                                31                      \n\n" +
				"     April 2024               May 2024               June 2024          \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				"    1  2  3  4  5  6              1  2  3  4                       1    \n" +
				" 7  8  9 10 11 12 13     5  6  7  8  9 10 11     2  3  4  5  6  7  8    \n" +
				"14 15 16 17 18 19 20    12 13 14 15 16 17 18     9 10 11 12 13 14 15    \n" +
				"21 22 23 24 25 26 27    19 20 21 22 23 24 25    16 17 18 19 20 21 22    \n" +
				"28 29 30                26 27 28 29 30 31       23 24 25 26 27 28 29    \n" +
				"                                                30                      \n\n" +
				"     July 2024              August 2024            September 2024       \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				"    1  2  3  4  5  6                 1  2  3     1  2  3  4  5  6  7    \n" +
				" 7  8  9 10 11 12 13     4  5  6  7  8  9 10     8  9 10 11 12 13 14    \n" +
				"14 15 16 17 18 19 20    11 12 13 14 15 16 17    15 16 17 18 19 20 21    \n" +
				"21 22 23 24 25 26 27    18 19 20 21 22 23 24    22 23 24 25 26 27 28    \n" +
				"28 29 30 31             25 26 27 28 29 30 31    29 30                   \n\n" +
				"    October 2024           November 2024           December 2024        \n" +
				"Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    Su Mo Tu We Th Fr Sa    \n" +
				"       1  2  3  4  5                    1  2     1  2  3  4  5  6  7    \n" +
				" 6  7  8  9 10 11 12     3  4  5  6  7  8  9     8  9 10 11 12 13 14    \n" +
				"13 14 15 16 17 18 19    10 11 12 13 14 15 16    15 16 17 18 19 20 21    \n" +
				"20 21 22 23 24 25 26    17 18 19 20 21 22 23    22 23 24 25 26 27 28    \n" +
				"27 28 29 30 31          24 25 26 27 28 29 30    29 30 31                \n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			stdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			DumpYear(tt.year)

			w.Close()
			os.Stdout = stdout
			if _, err := buf.ReadFrom(r); err != nil {
				t.Fatalf("Failed to read from pipe: %v", err)
			}

			output := stripAnsiCodes(buf.String())
			assert.Equal(t, tt.expected, output)
		})
	}
}
