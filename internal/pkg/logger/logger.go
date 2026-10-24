package logger

import (
	"context"
)

type Logger interface {
	Info(ctx context.Context, msg string, data ...any)
	Error(ctx context.Context, msg string, data ...any)
	Debug(ctx context.Context, msg string, data ...any)
	Warn(ctx context.Context, msg string, data ...any)
}
