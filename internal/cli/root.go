// Package cli wires the cal command-line interface to the calendar library.
package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/mojotx/cal/pkg/calendar"
	"github.com/spf13/cobra"
)

// nowFunc is overridden in tests to make "no arguments" behavior deterministic.
var nowFunc = time.Now

// NewRootCmd builds a new root command instance.
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cal [month] [year]",
		Short: "Print a calendar for a month, year, or the current month",
		Long: `cal prints a calendar to the terminal.

With no arguments, it prints the current month. With one argument, it
prints the entire calendar for that year. With two arguments, it prints
the calendar for the given month and year.`,
		Args:          cobra.MaximumNArgs(2),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          runRoot,
	}
	return cmd
}

// Execute runs the root command using the process's command-line arguments.
// Any error, whether from cobra itself (e.g. bad flags) or from runRoot, is printed here.
func Execute() error {
	err := NewRootCmd().Execute()
	if err != nil {
		color.Red("%s", err.Error())
	}
	return err
}

func runRoot(_ *cobra.Command, args []string) error {
	switch len(args) {
	case 0:
		now := nowFunc()
		calendar.DumpMonth(now.Month(), now.Year())
		return nil

	case 1:
		year, err := strconv.Atoi(args[0])
		if err != nil {
			return reportError("error parsing year: %s", err.Error())
		}
		if err := calendar.DumpYear(year); err != nil {
			return reportError("error: %s", err.Error())
		}
		return nil

	default: // 2 args
		month, err := strconv.Atoi(args[0])
		if err != nil {
			return reportError("error parsing month: %s", err.Error())
		}
		year, err := strconv.Atoi(args[1])
		if err != nil {
			return reportError("error parsing year: %s", err.Error())
		}
		if err := calendar.ValidateMonth(month); err != nil {
			return reportError("error: %s", err.Error())
		}
		if err := calendar.ValidateYear(year); err != nil {
			return reportError("error: %s", err.Error())
		}
		calendar.DumpMonth(time.Month(month), year)
		return nil
	}
}

// reportError builds an error for runRoot; Execute prints it so all errors share one output path.
func reportError(format string, args ...any) error {
	return fmt.Errorf(format, args...)
}
