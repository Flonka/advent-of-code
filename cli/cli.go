package cli

import (
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

// Default sets up default CLI args and slog package configuration.
func Default() *slog.LevelVar {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Enable debug logging")
	flag.Parse()

	logLevel := new(slog.LevelVar)
	if debug {
		logLevel.Set(slog.LevelDebug)
	}

	// Set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      logLevel,
			TimeFormat: time.StampMicro,
		}),
	))
	return logLevel
}
