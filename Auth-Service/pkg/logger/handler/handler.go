package handler

import (
	"Auth-Service/pkg/logger/data"
	"fmt"
	"log/slog"
	"path"
	"runtime"
	"strings"
)

var levelNames = map[slog.Level]string{
	data.LevelPanic: "panic",
	data.LevelFatal: "fatal",
}

// GetHandlerOptions  creates and returns log handler options
// customizing log level labels and source information.
func GetHandlerOptions(level string) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		Level: parserLevel(level),
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.LevelKey:
				l := a.Value.Any().(slog.Level)
				levelLabel, exists := levelNames[l]
				if exists {
					a.Value = slog.StringValue(levelLabel)
				}
			case slog.SourceKey:
				pc, file, line, _ := runtime.Caller(7)
				a.Value = slog.StringValue(fmt.Sprintf("%s [%s:%d]", path.Base(runtime.FuncForPC(pc).Name()), path.Base(file), line))
			}
			return a
		},
		AddSource: true,
	}
}

// parserLevel converts a string log level into its corresponding slog.Level value.
func parserLevel(level string) slog.Level {
	var lvl slog.Level

	switch strings.ToLower(level) {
	case "debug":
		lvl = slog.LevelDebug
	case "info":
		lvl = slog.LevelInfo
	case "warn", "warning":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	default:
		lvl = slog.LevelInfo
	}

	return lvl
}
