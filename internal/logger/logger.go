package logger

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Default() *slog.Logger {
	return logger
}
