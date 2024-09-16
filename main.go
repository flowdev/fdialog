package main

import (
	"github.com/flowdev/fdialog/cmd"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Remove time from the output
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}

			return a
		},
	}))
	slog.SetDefault(logger)
	cmd.Execute()
}
