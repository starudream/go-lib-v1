package main

import (
	"context"
	"fmt"

	"github.com/starudream/go-lib/app"
	"github.com/starudream/go-lib/log"
)

func main() {
	app.Add(wrap(fmt.Errorf("x")))
	err := app.OnceGo()
	if err != nil {
		log.Fatal().Msgf("app init fail: %v", err)
	}
}

func wrap(err error) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		return err
	}
}
