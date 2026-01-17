package handler

import (
	"log/slog"
	"testing"
)

func TestGetHandlerOptions(t *testing.T) {
	type args struct {
		level string
	}

	tests := []struct {
		name   string
		args   args
		assert func(t *testing.T, got *slog.HandlerOptions)
	}{
		{
			name: "Valid level 'DEBUG'",
			args: args{level: "DEBUG"},
			assert: func(t *testing.T, got *slog.HandlerOptions) {
				t.Helper()

				if got.Level != slog.LevelDebug {
					t.Errorf("Level = %v, want %v", got.Level, slog.LevelDebug)
				}

				if got.ReplaceAttr == nil {
					t.Error("ReplaceAttr should not be nil")
				}

				if !got.AddSource {
					t.Error("AddSource should be true")
				}
			},
		},
		{
			name: "Valid level 'info'",
			args: args{level: "info"},
			assert: func(t *testing.T, got *slog.HandlerOptions) {
				t.Helper()

				if got.Level != slog.LevelInfo {
					t.Errorf("Level = %v, want %v", got.Level, slog.LevelInfo)
				}
				if got.ReplaceAttr == nil {
					t.Error("ReplaceAttr should not be nil")
				}

				if !got.AddSource {
					t.Error("AddSource should be true")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetHandlerOptions(tt.args.level)
			tt.assert(t, got)
		})
	}
}

func Test_parserLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
		want slog.Level
	}{
		{
			name: "Valid level 'DEBUG'",
			args: args{level: "DEBUG"},
			want: slog.LevelDebug,
		},
		{
			name: "Default level for invalid input",
			args: args{level: "invalid"},
			want: slog.LevelInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parserLevel(tt.args.level); got != tt.want {
				t.Errorf("parserLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
