package log

import (
	"io"
	"os"
	"path/filepath"

	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/starudream/go-lib/constant"
)

func NewConsoleWriter() io.Writer {
	return &zerolog.ConsoleWriter{
		Out:        colorable.NewColorableStdout(),
		TimeFormat: constant.LoggerTimeFormat,
	}
}

type FileWriterConfig struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

func (c FileWriterConfig) init() FileWriterConfig {
	if c.Filename == "" {
		c.Filename = filepath.Join(os.TempDir(), filepath.Base(os.Args[0])+".log")
	}
	if c.MaxSize == 0 {
		c.MaxSize = 100
	}
	if c.MaxAge == 0 {
		c.MaxAge = 365
	}
	return c
}

func NewFileWriter(c FileWriterConfig) io.Writer {
	c = c.init()
	return &zerolog.ConsoleWriter{
		Out: &lumberjack.Logger{
			Filename:   c.Filename,
			MaxSize:    c.MaxSize,
			MaxAge:     c.MaxAge,
			MaxBackups: c.MaxBackups,
			LocalTime:  true,
		},
		NoColor:    true,
		TimeFormat: constant.LoggerTimeFormat,
	}
}
