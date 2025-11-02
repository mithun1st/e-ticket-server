package applogger

import (
	"log/slog"
	"os"
)

var customLogger *slog.Logger

func init() {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	customLogger = slog.New(handler)
}

func Info(data any) {
	customLogger.Info("", "data", data)
}

func Error(data any) {
	customLogger.Error("", "data", data)
}

// InfoWithMsg logs with both message and data
func InfoWithMsg(msg string, data any) {
	customLogger.Info(msg, "data", data)
}

// ErrorWithMsg logs with both error message and data
func ErrorWithMsg(msg string, data any) {
	customLogger.Error(msg, "data", data)
}
