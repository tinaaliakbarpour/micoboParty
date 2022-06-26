package logger

import (
	"fmt"
	"micobianParty/config"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Add new log
func (l *logs) Add(key string, field interface{}) *logs {
	l.fields = append(l.fields, zap.Any(key, field))
	return l
}

// Append new log
func (l *logs) Append(fields ...zapcore.Field) *logs {
	l.fields = append(l.fields, fields...)
	return l
}

// Commit meth
func (l *logs) Commit(message string) {
	l.Add("service", config.Confs.Service.Name)

	defer func() {
		l.logger.Sync()
		l.fields = nil
	}()

	switch l.level {
	case zapcore.InfoLevel:
		l.logger.Info(message, l.fields...)
	case zapcore.WarnLevel:
		l.logger.Warn(message, l.fields...)
	case zapcore.DebugLevel:
		l.logger.Debug(message, l.fields...)
	case zapcore.FatalLevel:
		l.logger.DPanic(message, l.fields...)
	default:
		l.logger.Warn(message, l.fields...)
	}
}

// Level of log
func (l logs) Level(level zapcore.Level) *logs {
	l.level = level
	return &l
}

// Development method
func (l *logs) Development() *logs {
	var caller string = ""

	pc, _, line, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		caller = details.Name()
	}

	l.Add("line", line)
	l.Add("caller", caller)

	return l
}

// Prepare new logger
func Prepare(zap *zap.Logger) *logs {
	return &logs{logger: zap}
}

// GetZapLogger func
// create new zap logger
func GetZapLogger(debug bool) *zap.Logger {
	once.Do(func() {
		var err error
		logger, err = zap.NewProduction()
		if err != nil {
			fmt.Errorf("Can't initialize the Logger with error %v", err)
		}
		defer logger.Sync()
	})

	return logger
}
