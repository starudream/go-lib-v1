package log

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/starudream/go-lib/internal/ilog"
)

type L = zerolog.Logger

var _l zerolog.Logger

func Logger() zerolog.Logger {
	return _l
}

func With() zerolog.Context {
	return _l.With()
}

func WithLevel(level zerolog.Level) *zerolog.Event {
	return _l.WithLevel(level)
}

func Trace() *zerolog.Event {
	return _l.Trace()
}

func Debug() *zerolog.Event {
	return _l.Debug()
}

func Info() *zerolog.Event {
	return _l.Info()
}

func Warn() *zerolog.Event {
	return _l.Warn()
}

func Error() *zerolog.Event {
	return _l.Error()
}

func Fatal() *zerolog.Event {
	return _l.Fatal()
}

func Panic() *zerolog.Event {
	return _l.Panic()
}

func Log() *zerolog.Event {
	return _l.Log()
}

func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

func Attach(kvs ...string) {
	vs := make([]any, len(kvs))
	for i := 0; i < len(kvs); i++ {
		vs[i] = kvs[i]
	}
	l := _l.With().Fields(vs).Logger()
	SetLogger(l)
}

func SetLogger(l zerolog.Logger) {
	_l = l
	ilog.X = _l
	log.Logger = _l
	zerolog.DefaultContextLogger = &_l
}
