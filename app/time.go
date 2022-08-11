package app

import (
	"time"
)

var (
	startup = time.Now()

	running time.Time
)

func StartupTime() time.Time {
	return startup
}

func RunningTime() time.Time {
	return running
}

func CostTime() time.Duration {
	return running.Sub(startup)
}
