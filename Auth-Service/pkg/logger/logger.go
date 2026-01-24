package logger

import (
	"Auth-Service/pkg/logger/console"
	LogHandler "Auth-Service/pkg/logger/handler"
	"context"
	"log/slog"
	"os"
)

type ctxKey string

const traceIDKey ctxKey = "trace_id"

type logger struct {
	log *slog.Logger
	opt *console.Options
}

func NewLogger(opts ...*console.Options) ILogger {
	var log *slog.Logger
	opt := buildOpts(opts...)

	switch opt.Mode {
	case console.ModeJSON:
		log = slog.New(slog.NewJSONHandler(os.Stdout, LogHandler.GetHandlerOptions(opt.Level)))
	case console.ModeText:
		log = slog.New(slog.NewTextHandler(os.Stdout, LogHandler.GetHandlerOptions(opt.Level)))
	default:
		log = slog.New(slog.NewJSONHandler(os.Stdout, LogHandler.GetHandlerOptions(opt.Level)))
	}

	return &logger{
		log: log,
		opt: opt,
	}
}

// Info logs an informational message with context.
func (l *logger) Info(ctx context.Context, msg string, keysAndValues ...any) {
	l.log.With(l.attrsFromCtx(ctx)...).Log(ctx, slog.LevelInfo, msg, keysAndValues...)
}

func (l *logger) Error(ctx context.Context, msg string, keysAndValues ...any) {
	l.log.With(l.attrsFromCtx(ctx)...).Log(ctx, slog.LevelError, msg, keysAndValues...)
}

// withTraceID adds a trace ID to the context for logging.
func (l *logger) withTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// attrsFromCtx extracts attributes from the context for logging.
func (l *logger) attrsFromCtx(ctx context.Context) []any {
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

func buildOpts(options ...*console.Options) *console.Options {
	if len(options) > 0 {
		return options[0]
	}
	return &console.Options{}
}
