package main

import (
	"os"

	"github.com/mojotx/cal/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
