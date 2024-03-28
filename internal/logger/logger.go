package logger

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	logger.Info("slog logger is created")
	return logger
}
