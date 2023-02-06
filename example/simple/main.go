package main

import (
	"context"
	"time"

	"github.com/starudream/go-lib/app"
	"github.com/starudream/go-lib/cache"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/httpx"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/randx"
	"github.com/starudream/go-lib/seq"
	"github.com/starudream/go-lib/timex"
)

func main() {
	log.Attach("app", "example-simple")
	app.Init(func() error { log.Info().Msg("init"); return nil })
	app.Add(wrapError(TestAppTime))
	app.Add(wrapError(TestCache))
	app.Add(wrapError(TestConfig))
	app.Add(wrapError(TestHTTPX))
	app.Add(wrapError(TestRandX))
	app.Add(wrapError(TestSeq))
	app.Add(wrapError(TestTimeX))
	app.Defer(TestDefer)
	err := app.OnceGo()
	if err != nil {
		panic(err)
	}
	log.Info().Msg("success")
}

func wrapError(f func()) func(ctx context.Context) error {
	return func(ctx context.Context) error { f(); return nil }
}

func TestAppTime() {
	log.Info().Msgf("startup: %v", app.StartupTime().Format(time.RFC3339Nano))
	log.Info().Msgf("running: %v", app.RunningTime().Format(time.RFC3339Nano))
}

func TestCache() {
	c := cache.SIMPLE()
	err := c.Set("foo", true)
	if err != nil {
		log.Fatal().Msgf("set foo error: %v", err)
	}
	v1, err := c.Get("foo")
	if err != nil {
		log.Fatal().Msgf("get foo error: %v", err)
	}
	log.Info().Msgf("foo: %v", v1)
	v2, err := c.Get("bar")
	if err != nil && !cache.IsErrKeyNotFound(err) {
		log.Fatal().Msgf("get bar error: %v", err)
	}
	log.Info().Msgf("bar: %v", v2)
}

func TestConfig() {
	log.Warn().Msgf("debug: %v", config.GetBool("debug"))
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

func TestTimeX() {
	log.Info().Msgf("begin of day: %s", timex.BeginOfDay().Format(timex.DateTimeFormat))
}

func TestDefer() {
	log.Info().Msg("bye bye")
}
