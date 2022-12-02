package ilog

import (
	"os"

	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/spf13/cast"

	"github.com/starudream/go-lib/constant"
)

var X = func() zerolog.Logger {
	l := func() zerolog.Level {
		if cast.ToBool(os.Getenv("DEBUG")) || (constant.PREFIX != "" && cast.ToBool(os.Getenv(constant.PREFIX+"_DEBUG"))) {
			return zerolog.DebugLevel
		}
		for i, a := range os.Args {
			if a == "--debug" && (i+1 >= len(os.Args) || (i+1 < len(os.Args) && cast.ToBool(os.Args[i+1]))) {
				return zerolog.DebugLevel
			}
		}
		return zerolog.InfoLevel
	}()
	lc := zerolog.New(&zerolog.ConsoleWriter{Out: colorable.NewColorableStdout(), TimeFormat: constant.LoggerTimeFormat}).Level(l).With()
	if constant.VERSION != "" {
		lc = lc.Str("version", constant.VERSION)
	}
	return lc.Timestamp().Logger()
}()
