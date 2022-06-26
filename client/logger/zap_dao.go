package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once   sync.Once
	logger *zap.Logger
)

// Log interface
type Log interface {
	Add(key string, field interface{}) *Log
	Append(fields ...zapcore.Field) *Log
	Level(level zapcore.Level) *Log
	Development() *Log
}

type logs struct {
	logger *zap.Logger
	fields []zapcore.Field
	level  zapcore.Level
}
