package main

import (
	"context"

	"github.com/starudream/go-lib/app"
	"github.com/starudream/go-lib/config"
	"github.com/starudream/go-lib/httpx"
	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/randx"
)

func main() {
	log.Attach("app", "example-simple")
	app.Add(wrapError(TestAppTime))
	app.Add(wrapError(TestConfig))
	// app.Add(wrapError(TestHTTPX))
	app.Add(wrapError(TestRandX))
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
	log.Info().Msgf("startup: %v", app.StartupTime())
	log.Info().Msgf("running: %v", app.RunningTime())
	log.Info().Msgf("cost: %v", app.CostTime())
}

func TestConfig() {
	log.Warn().Msgf("debug: %v", config.GetBool("debug"))
}

func TestHTTPX() {
	resp, err := httpx.R().Get("https://api.github.com")
	if err != nil {
		panic(err)
	}
	log.Debug().Msgf("response status: %s", resp.Status())
}

func TestRandX() {
	log.Info().Msgf("fake %s", randx.F().LetterN(16))
}

func TestDefer() {
	log.Info().Msg("bye bye")
}
