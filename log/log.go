package log

import (
	"io"

	"github.com/rs/zerolog"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/internal/ilog"
)

func init() {
	zerolog.CallerSkipFrameCount = 2
	zerolog.InterfaceMarshalFunc = json.Marshal
	zerolog.TimeFieldFormat = constant.LoggerTimeFormat

	var (
		level   zerolog.Level
		writers []io.Writer
	)

	if config.GetBool("debug") {
		level = zerolog.DebugLevel
	} else if s := config.GetString("log.level"); s != "" {
		lvl, err := zerolog.ParseLevel(s)
		if err != nil {
			ilog.X.Fatal().Msgf("invalid log level: %s", s)
		}
		level = lvl
	}

	if !config.IsSet("log.console") || config.GetBool("log.console") {
		writers = append(writers, NewConsoleWriter())
	}

	if p := config.GetString("log.file.path"); p != "" {
		writers = append(writers, NewFileWriter(FileWriterConfig{
			Filename:   p,
			MaxSize:    config.GetInt("log.file.max_size"),
			MaxAge:     config.GetInt("log.file.max_age"),
			MaxBackups: config.GetInt("log.file.max_backups"),
		}))
	}

	l := func() zerolog.Logger {
		if len(writers) == 0 {
			return zerolog.Nop()
		}
		lc := zerolog.New(zerolog.MultiLevelWriter(writers...)).Level(level).With()
		if config.GetBool("debug") {
			lc = lc.Caller()
		}
		if constant.VERSION != "" {
			lc = lc.Str("version", constant.VERSION)
		}
		return lc.Timestamp().Logger()
	}()

	SetLogger(l)
}
