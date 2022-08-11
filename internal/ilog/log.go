package ilog

import (
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"

	"github.com/starudream/go-lib/constant"
)

var X = func() zerolog.Logger {
	lc := zerolog.New(&zerolog.ConsoleWriter{Out: colorable.NewColorableStdout(), TimeFormat: constant.LoggerTimeFormat}).With()
	if constant.VERSION != "" {
		lc = lc.Str("version", constant.VERSION)
	}
	return lc.Timestamp().Logger()
}()
