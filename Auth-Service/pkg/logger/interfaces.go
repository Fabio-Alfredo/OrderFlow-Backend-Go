package logger

import (
	"context"
)

type ILogger interface {
	Info(ctx context.Context, msg string, keysAndValues ...any)
	Error(ctx context.Context, msg string, keysAndValues ...any)
}
