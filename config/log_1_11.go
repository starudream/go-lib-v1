package config

import (
	"fmt"

	"github.com/starudream/go-lib/internal/ilog"
	"github.com/starudream/go-lib/internal/viper"
)

type logger struct {
}

var _ viper.Logger = (*logger)(nil)

func (l *logger) Trace(msg string, kvs ...any) {
	ilog.X.Trace().CallerSkipFrame(1).Msg(concat(msg, kvs...))
}

func (l *logger) Debug(msg string, kvs ...any) {
	ilog.X.Debug().CallerSkipFrame(1).Msg(concat(msg, kvs...))
}

func (l *logger) Info(msg string, kvs ...any) {
	ilog.X.Debug().CallerSkipFrame(1).Msg(concat(msg, kvs...))
}

func (l *logger) Warn(msg string, kvs ...any) {
	ilog.X.Warn().CallerSkipFrame(1).Msg(concat(msg, kvs...))
}

func (l *logger) Error(msg string, kvs ...any) {
	ilog.X.Error().CallerSkipFrame(1).Msg(concat(msg, kvs...))
}

func concat(msg string, kvs ...any) string {
	out := msg

	if len(kvs) > 0 && len(kvs)%2 == 1 {
		kvs = append(kvs, nil)
	}

	for i := 0; i <= len(kvs)-2; i += 2 {
		out = fmt.Sprintf("%s %v=%v", out, kvs[i], kvs[i+1])
	}

	return out
}
