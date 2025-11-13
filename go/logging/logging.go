package logging

import (
	"context"
	"log/slog"
	"os"
	"sync"
)

var (
	mu              sync.RWMutex
	defaultLogger   = slog.Default()
	slogLogLevelMap = map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}
)

type ctxKey struct{}

// loggerを埋め込んだctxを返す
// interceptorなどで使うことを想定している
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

// ctxからloggerを取り出す
// WithLoggerですでにctxにはloggerが入っていることを想定している
func FromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*slog.Logger); ok && l != nil {
		return l
	}
	return defaultLogger
}

func New(logLevel string, attrs ...slog.Attr) *slog.Logger {
	handlerAttrs := []slog.Attr{}
	for _, attr := range attrs {
		if attr.Key != "" {
			handlerAttrs = append(handlerAttrs, attr)
		}
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slogLogLevelMap[logLevel],
	}).WithAttrs(handlerAttrs))
	return logger
}

func SetDefaultLogger(logger *slog.Logger) {
	mu.Lock()
	defer mu.Unlock()
	defaultLogger = logger
}
