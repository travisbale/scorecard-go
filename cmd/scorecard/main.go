package main

import (
	"log/slog"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:     "scorecard",
		Usage:    "Score keeping API for the Manitoba Ryder Cup",
		Flags:    []cli.Flag{},
		Commands: []*cli.Command{},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("server shutdown", err, err)
	}
}
