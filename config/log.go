//go:build !viper_logger

package config

import (
	"bytes"
	"io"
	"log"

	"github.com/rs/zerolog"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"

	"github.com/starudream/go-lib/internal/ilog"
)

func New(opts ...viper.Option) *viper.Viper {
	initLogger()
	return viper.NewWithOptions(opts...)
}

func initLogger() {
	jww.TRACE = newLogger(zerolog.TraceLevel)
	jww.DEBUG = newLogger(zerolog.DebugLevel)
	jww.INFO = newLogger(zerolog.DebugLevel)
	jww.WARN = newLogger(zerolog.WarnLevel)
	jww.ERROR = newLogger(zerolog.ErrorLevel)
	jww.CRITICAL = newLogger(zerolog.ErrorLevel)
	jww.FATAL = newLogger(zerolog.ErrorLevel)
}

func newLogger(level zerolog.Level) *log.Logger {
	return log.New(&writer{level: level}, "", 0)
}

type writer struct {
	level zerolog.Level
}

var _ io.Writer = (*writer)(nil)

func (w *writer) Write(p []byte) (n int, err error) {
	ilog.X.WithLevel(w.level).Msg(string(bytes.TrimSuffix(p, []byte{'\n'})))
	return len(p), nil
}
