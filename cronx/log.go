package cronx

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
)

type logger struct {
	Logger zerolog.Logger
}

var _ cron.Logger = (*logger)(nil)

func (l *logger) Info(msg string, keysAndValues ...any) {
	l.Logger.Info().Msgf(splice(msg, keysAndValues...))
}

func (l *logger) Error(err error, msg string, keysAndValues ...any) {
	l.Logger.Err(err).Msgf(splice(msg, keysAndValues...))
}

func splice(msg string, kvs ...any) string {
	out := msg

	if len(kvs) > 0 && len(kvs)%2 == 1 {
		kvs = append(kvs, nil)
	}

	for i := 0; i <= len(kvs)-2; i += 2 {
		out = fmt.Sprintf("%s %v=%v", out, kvs[i], kvs[i+1])
	}

	return out
}
