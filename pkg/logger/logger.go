package logger

import (
	"io"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger(writer io.Writer) Logger {
	return Logger{logger: zerolog.New(writer)}
}

func NewNopLogger() Logger {
	return Logger{logger: zerolog.Nop()}
}

func (l Logger) Info(key string, fields map[string]interface{}) {
	l.logger.Info().Fields(fields).Msg(key)
}

func (l Logger) Error(key string, fields map[string]interface{}) {
	l.logger.Error().Fields(fields).Msg(key)
}
