package console

import "log/slog"

const (
	ModeJSON    = "json"
	ModeText    = "text"
	LevelFatal  = slog.Level(13)
	LevelPanic  = slog.Level(12)
	ErrorKey    = "error"
	RequestKey  = "request"
	ResponseKey = "response"
	StartKey    = "start"
	EndKey      = "finish"
	FailedKey   = "failed"
	DataKey     = "data"
)
