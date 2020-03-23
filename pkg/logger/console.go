package logger

import (
	"context"
	"log"
	"os"
)

const (
	CLR_R = "\x1b[31;1m"
	CLR_0 = "\x1b[33;1m"
	CLR_B = "\x1b[34;1m"
	CLR_W = "\x1b[37;1m"
	CLR_Z = "\x1b[0m"
)

func init() {
	factoryFunc = NewConsoleLogger
}

type consoleLogger struct {
	level  uint8
	logger *log.Logger
}

func NewConsoleLogger(level uint8) Logger {
	return &consoleLogger{
		level:  level,
		logger: log.New(os.Stdout, "", DefaultFlags),
	}
}

func (c consoleLogger) Debug(ctx context.Context, format string, args ...interface{}) {
	if c.level|DebugLevel != 0 {
		c.logger.Printf("%s: "+format, append([]interface{}{CLR_W + DebugPrefix + CLR_Z}, args...)...)
	}
}

func (c consoleLogger) Info(ctx context.Context, format string, args ...interface{}) {
	if c.level&InfoLevel != 0 {
		c.logger.Printf("%s: "+format, append([]interface{}{CLR_B + InfoPrefix + CLR_Z}, args...)...)
	}
}

func (c consoleLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	if c.level&WarnLevel != 0 {
		c.logger.Printf("%s: "+format, append([]interface{}{CLR_0 + WarnPrefix + CLR_Z}, args...)...)
	}
}

func (c consoleLogger) Error(ctx context.Context, format string, args ...interface{}) {
	if c.level&ErrorLevel != 0 {
		c.logger.Printf("%s: "+format, append([]interface{}{CLR_R + ErrorPrefix + CLR_Z}, args...)...)
	}
}

func (c consoleLogger) Fatal(ctx context.Context, format string, args ...interface{}) {
	if c.level&FatalLevel != 0 {
		c.logger.Printf("%s: "+format, append([]interface{}{CLR_R + FatalPrefix + CLR_Z}, args...)...)
	}
}

func (c *consoleLogger) SetLevel(level uint8) {
	c.level = level
}