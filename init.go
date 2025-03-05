package logger

import (
	"log/slog"
	"os"
)

// Logger instance used by external software.
var logger *slog.Logger

// Setting up default logger in development conf
func init() {
	err := Init(Config{})
	if err != nil {
		panic(err)
	}
}

// Init ...
func Init(c Config) error {
	var level slog.Level
	switch c.Mode {
	case Staging:
		level = slog.LevelDebug
	case Production:
		level = slog.LevelInfo
	default:
		level = slog.LevelDebug
	}

	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
	logger = slog.New(h)

	return nil
}
