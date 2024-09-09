package log

import (
	"log/slog"
	"os"
)

func NewLoggerService() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}

func SetLoggerLevel(level string) *slog.Logger {
	var logLevel slog.Level
	if level == "warn" || level == "Warn" || level == "WARN" {
		logLevel = 4
	} else if level == "error" || level == "Error" || level == "ERROR" {
		logLevel = 8
	} else if level == "info" || level == "Info" || level == "INFO" {
		logLevel = 0
	} else {
		logLevel = -4
	}
	opts1 := &slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts1))
	return logger
}
