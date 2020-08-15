package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func DevLogger() *Logger {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &Logger{
		Logger: l,
	}
}

func init() {
	Logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(Logger)
}
