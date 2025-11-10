package logging

import (
	"log/slog"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
)

func NewDDTraceLogger(logger *slog.Logger) ddtrace.Logger {
	return &ddTraceLogger{logger}
}

type ddTraceLogger struct {
	logger *slog.Logger
}

func (l *ddTraceLogger) Log(msg string) {
	l.logger.Info(msg)
}

// ddTraceLoggerはddtrace.Logger interfaceを満たしている
var _ ddtrace.Logger = (*ddTraceLogger)(nil)
