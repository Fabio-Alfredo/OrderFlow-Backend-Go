package logger

import (
	"Auth-Service/pkg/logger/data"
	"Auth-Service/pkg/logger/handler"
	"bytes"
	"context"
	"log/slog"
	"os"
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		opts *data.Options
	}

	tests := []struct {
		name   string
		args   args
		assert func(t *testing.T, got ILogger)
	}{
		{
			name: "JSON mode with DEBUG level",
			args: args{
				opts: &data.Options{
					Mode:  data.ModeJSON,
					Level: "DEBUG",
				},
			},
			assert: func(t *testing.T, got ILogger) {
				t.Helper()

				l, ok := got.(*logger)
				if !ok {
					t.Fatalf("expected *logger, got %T", got)
				}

				if l.opt.Mode != data.ModeJSON {
					t.Errorf("Mode = %v, want %v", l.opt.Mode, data.ModeJSON)
				}

				if l.opt.Level != "DEBUG" {
					t.Errorf("Level = %v, want DEBUG", l.opt.Level)
				}

				if l.log == nil {
					t.Fatal("logger.log should not be nil")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLogger(tt.args.opts)
			tt.assert(t, got)
		})
	}
}

func Test_logger_Info(t *testing.T) {
	type fields struct {
		log *slog.Logger
		opt *data.Options
	}
	type args struct {
		ctx           context.Context
		msg           string
		keysAndValues []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Log info message without context",
			fields: fields{
				log: slog.New(slog.NewTextHandler(os.Stdout, handler.GetHandlerOptions("INFO"))),
				opt: &data.Options{
					Mode:  data.ModeText,
					Level: "INFO",
				},
			},
			args: args{
				ctx: context.Background(),
				msg: "This is an info message",
			},
		},
		{
			name: "Log info message with trace ID in context",
			fields: fields{
				log: slog.New(slog.NewTextHandler(os.Stdout, handler.GetHandlerOptions("INFO"))),
				opt: &data.Options{
					Mode:  data.ModeText,
					Level: "INFO",
				},
			},
			args: args{
				ctx: func() context.Context {
					l := &logger{}
					return l.withTraceID(context.Background(), "trace-12345")
				}(),
				msg: "This is an info message with trace ID",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logger{
				log: tt.fields.log,
				opt: tt.fields.opt,
			}
			l.Info(tt.args.ctx, tt.args.msg, tt.args.keysAndValues...)
		})
	}
}

func Test_logger_withTraceID(t *testing.T) {
	var buf bytes.Buffer

	type fields struct {
		log *slog.Logger
		opt *data.Options
	}
	type args struct {
		ctx     context.Context
		traceID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   context.Context
	}{
		{
			name: "Log info message without context",
			fields: fields{
				log: slog.New(slog.NewTextHandler(&buf, handler.GetHandlerOptions("INFO"))),
				opt: &data.Options{
					Mode:  data.ModeText,
					Level: "INFO",
				},
			},
			args: args{
				ctx:     context.Background(),
				traceID: "trace-12345",
			},
			want: func() context.Context {
				return context.WithValue(context.Background(), traceIDKey, "trace-12345")
			}(),
		},
		{
			name: "Log info message with existing context",
			fields: fields{
				log: slog.New(slog.NewTextHandler(&buf, handler.GetHandlerOptions("INFO"))),
				opt: &data.Options{
					Mode:  data.ModeText,
					Level: "INFO",
				},
			},
			args: args{
				ctx:     context.WithValue(context.Background(), "existing_key", "existing_value"),
				traceID: "trace-67890",
			},
			want: func() context.Context {
				return context.WithValue(context.WithValue(context.Background(), "existing_key", "existing_value"), traceIDKey, "trace-67890")
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &logger{
				log: tt.fields.log,
				opt: tt.fields.opt,
			}
			if got := l.withTraceID(tt.args.ctx, tt.args.traceID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("withTraceID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_attrsFromCtx(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "Context with trace ID",
			args: args{
				ctx: func() context.Context {
					l := &logger{}
					return l.withTraceID(context.Background(), "trace-12345")
				}(),
			},
			want: []any{
				slog.String("trace_id", "trace-12345"),
			},
		},
		{
			name: "Context without trace ID",
			args: args{
				ctx: context.Background(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := attrsFromCtx(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("attrsFromCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
