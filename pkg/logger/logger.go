package logger

import (
	"context"
	"log"
)

type Logger interface {
	Debug(ctx context.Context, format string, args ...interface{})
	Info(ctx context.Context, format string, args ...interface{})
	Warn(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, format string, args ...interface{})
	Fatal(ctx context.Context, format string, args ...interface{})
	SetLevel(level uint8)
}

const (
	FatalLevel uint8 = 1 << iota
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel

	FatalPrefix = "FATAL"
	ErrorPrefix = "ERROR"
	WarnPrefix  = "WARN"
	InfoPrefix  = "INFO"
	DebugPrefix = "DEBUG"

	DefaultFlags = log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC
	DefaultLevel = FatalLevel | ErrorLevel | WarnLevel | InfoLevel | DebugLevel
)

type factory func(level uint8) Logger

var factoryFunc factory

func NewGlobalLogger(level uint8) {
	logger = factoryFunc(level)
}

var logger Logger

func Debug(ctx context.Context, format string, args ...interface{}) {
	logger.Debug(ctx, format, args...)
}
func Info(ctx context.Context, format string, args ...interface{}) {
	logger.Info(ctx, format, args...)
}
func Warn(ctx context.Context, format string, args ...interface{}) {
	logger.Warn(ctx, format, args...)
}
func Error(ctx context.Context, format string, args ...interface{}) {
	logger.Error(ctx, format, args...)
}
func Fatal(ctx context.Context, format string, args ...interface{}) {
	logger.Fatal(ctx, format, args...)
}
func SetLevel(level uint8) {
	logger.SetLevel(level)
}
