package httpx

import (
	"github.com/go-resty/resty/v2"

	"github.com/starudream/go-lib/log"
)

type logger struct {
	Logger log.L
}

var _ resty.Logger = (*logger)(nil)

func (l *logger) Errorf(format string, v ...any) {
	l.Logger.Error().Msgf(format, v...)
}

func (l *logger) Warnf(format string, v ...any) {
	l.Logger.Warn().Msgf(format, v...)
}

func (l *logger) Debugf(format string, v ...any) {
	l.Logger.Debug().Msgf(format, v...)
}
