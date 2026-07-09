package eslog

import (
	"context"
	"log/slog"
)

func Noisy(msg string, args ...any) {
	slog.Log(context.Background(), -8, msg, args...)
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}
func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}
func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}
func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}
func Crit(msg string, args ...any) {
	slog.Log(context.Background(), 12, msg, args...)
}
