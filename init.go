package logger

import (
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/log/global"
)

// Logger instance used by external software.
var logger *slog.Logger

func setup(c Config) error {
	var level slog.Level
	switch c.Mode {
	case Staging:
		level = slog.LevelDebug
	case Production:
		level = slog.LevelInfo
	default:
		level = slog.LevelDebug
	}

	logger = slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: level,
			}),
			otelslog.NewHandler("otel", otelslog.WithLoggerProvider(global.GetLoggerProvider())),
		),
	)

	return nil
}
