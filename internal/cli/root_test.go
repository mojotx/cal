package cli

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// captureStdout redirects os.Stdout for the duration of fn and returns what was written.
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	stdout := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w

	fn()

	require.NoError(t, w.Close())
	os.Stdout = stdout

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	require.NoError(t, err)
	return buf.String()
}

func TestRunRoot_NoArgs(t *testing.T) {
	originalNow := nowFunc
	nowFunc = func() time.Time {
		return time.Date(2025, time.July, 15, 12, 0, 0, 0, time.UTC)
	}
	defer func() { nowFunc = originalNow }()

	cmd := NewRootCmd()
	cmd.SetArgs([]string{})
	out := captureStdout(t, func() {
		assert.NoError(t, cmd.Execute())
	})
	assert.Contains(t, out, "July 2025")
}

func TestRunRoot_YearOnly(t *testing.T) {
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"2024"})
	out := captureStdout(t, func() {
		assert.NoError(t, cmd.Execute())
	})
	assert.Contains(t, out, "February 2024")
}

func TestRunRoot_MonthAndYear(t *testing.T) {
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"2", "2024"})
	out := captureStdout(t, func() {
		assert.NoError(t, cmd.Execute())
	})
	assert.Contains(t, out, "February 2024")
}

func TestRunRoot_InvalidYearArg(t *testing.T) {
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"not-a-year"})
	_ = captureStdout(t, func() {
		err := cmd.Execute()
		assert.ErrorContains(t, err, "error parsing year")
	})
}

func TestRunRoot_InvalidMonthArg(t *testing.T) {
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"not-a-month", "2024"})
	_ = captureStdout(t, func() {
		err := cmd.Execute()
		assert.ErrorContains(t, err, "error parsing month")
	})
}

func TestRunRoot_OutOfRangeMonth(t *testing.T) {
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"13", "2024"})
	_ = captureStdout(t, func() {
		err := cmd.Execute()
		assert.ErrorContains(t, err, "month must be between 1 and 12")
	})
}

func TestRunRoot_OutOfRangeYear(t *testing.T) {
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"1", "0"})
	_ = captureStdout(t, func() {
		err := cmd.Execute()
		assert.ErrorContains(t, err, "year must be greater than 0")
	})
}

func TestRunRoot_TooManyArgs(t *testing.T) {
	cmd := NewRootCmd()
	cmd.SetArgs([]string{"1", "2024", "extra"})
	_ = captureStdout(t, func() {
		err := cmd.Execute()
		assert.ErrorContains(t, err, "accepts at most 2 arg(s)")
	})
}

func TestExecute(t *testing.T) {
	originalArgs := os.Args
	os.Args = []string{"cal", "2024"}
	defer func() { os.Args = originalArgs }()

	out := captureStdout(t, func() {
		assert.NoError(t, Execute())
	})
	assert.Contains(t, out, "February 2024")
}
