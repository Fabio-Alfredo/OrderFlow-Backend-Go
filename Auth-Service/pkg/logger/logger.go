package logger

import (
	"Auth-Service/pkg/logger/data"
	LogHandler "Auth-Service/pkg/logger/handler"
	"context"
	"log/slog"
	"os"
)

type ctxKey string

const traceIDKey ctxKey = "trace_id"

type logger struct {
	log *slog.Logger
	opt *data.Options
}

func NewLogger(opts *data.Options) ILogger {
	var log *slog.Logger

	switch opts.Mode {
	case data.ModeJSON:
		log = slog.New(slog.NewJSONHandler(os.Stdout, LogHandler.GetHandlerOptions(opts.Level)))
	case data.ModeText:
		log = slog.New(slog.NewTextHandler(os.Stdout, LogHandler.GetHandlerOptions(opts.Level)))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, LogHandler.GetHandlerOptions(opts.Level)))
	}

	return &logger{
		log: log,
		opt: opts,
	}
}

// Info logs an informational message with context.
func (l *logger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	l.log.With(attrsFromCtx(ctx)...).Log(ctx, slog.LevelInfo, msg, keysAndValues...)
}

// withTraceID adds a trace ID to the context for logging.
func (l *logger) withTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// attrsFromCtx extracts attributes from the context for logging.
func attrsFromCtx(ctx context.Context) []any {
	if ctx == nil {
		return nil
	}

	if v, ok := ctx.Value(traceIDKey).(string); ok {
		return []any{
			slog.String("trace_id", v),
		}
	}

	return nil
}
