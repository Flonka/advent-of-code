package cli

import (
	"flag"
	"log/slog"
	"os"
)

// Default sets up default CLI args and slog package configuration.
func Default() *slog.LevelVar {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Enable debug logging")
	flag.Parse()

	logLevel := new(slog.LevelVar)
	if debug == true {
		logLevel.Set(slog.LevelDebug)
	}

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(h))
	return logLevel

}
