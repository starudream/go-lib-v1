package log

import (
	"io"
	"strconv"
	"strings"

	"github.com/rs/zerolog"

	"github.com/starudream/go-lib/codec/json"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/constant"
	"github.com/starudream/go-lib/internal/ilog"
)

func init() {
	zerolog.CallerSkipFrameCount = 2
	zerolog.CallerMarshalFunc = customCallerMarshalFunc
	zerolog.InterfaceMarshalFunc = json.Marshal
	zerolog.TimeFieldFormat = constant.LoggerTimeFormat

	var (
		level   = zerolog.InfoLevel
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

func customCallerMarshalFunc(_ uintptr, file string, line int) string {
	fs := strings.Split(file, "/")
	idx := func() int {
		if len(fs) < 2 {
			return 0
		}
		for i := 0; i < len(fs); i++ {
			if strings.Contains(fs[i], "@") {
				if i > 0 {
					return i - 1
				}
				return 0
			}
		}
		return len(fs) - 2
	}()
	return strings.Join(fs[idx:], "/") + ":" + strconv.Itoa(line)
}
