package main

import (
	"context"
	"time"

	"github.com/starudream/go-lib/app"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/errx"
	"github.com/starudream/go-lib/httpx"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/randx"
	"github.com/starudream/go-lib/seq"
)

func main() {
	log.Attach("app", "example-simple")
	app.Add(wrapError(TestAppTime))
	app.Add(wrapError(TestConfig))
	app.Add(wrapError(TestErrx))
	app.Add(wrapError(TestHTTPX))
	app.Add(wrapError(TestRandX))
	app.Add(wrapError(TestSeq))
	app.Defer(TestDefer)
	err := app.OnceGo()
	if err != nil {
		panic(err)
	}
}

func wrapError(f func()) func(ctx context.Context) error {
	return func(ctx context.Context) error { f(); return nil }
}

func TestAppTime() {
	log.Info().Msgf("startup: %v", app.StartupTime().Format(time.RFC3339Nano))
	log.Info().Msgf("running: %v", app.RunningTime().Format(time.RFC3339Nano))
	log.Info().Msgf("cost: %v", app.CostTime())
}

func TestConfig() {
	log.Warn().Msgf("debug: %v", config.GetBool("debug"))
}

func TestErrx() {
	e1 := errx.New("a1")
	e2 := errx.New("b1")
	e3 := errx.Wrap(e1, "a11")
	log.Info().Msgf("%#v", e1)
	log.Info().Msgf("%#v", e2)
	log.Info().Msgf("%+v", e3)
	e4 := errx.Cause(e3)
	log.Info().Msgf("%#v", e4)
	if !errx.Is(e1, e4) {
		panic("e1 == e4")
	}
}

func TestHTTPX() {
	httpx.SetTimeout(3 * time.Second)
	httpx.SetUserAgent("go")
	resp, err := httpx.R().Get("https://www.gstatic.com/generate_204")
	if err != nil {
		panic(err)
	}
	log.Debug().Msgf("response status: %s", resp.Status())
}

func TestRandX() {
	log.Info().Msgf("fake: %s", randx.F().LetterN(16))
}

func TestSeq() {
	log.Info().Msgf("sonyflake: %s", seq.NextId())
	log.Info().Msgf("uuid: %s", seq.UUID())
	log.Info().Msgf("uuid short: %s", seq.UUIDShort())
}

func TestDefer() {
	log.Info().Msg("bye bye")
}
