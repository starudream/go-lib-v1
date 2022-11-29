package app

import (
	"time"
)

var (
	startup = time.Now()

	initing time.Time
	running time.Time
)

func StartupTime() time.Time {
	return startup
}

func InitingTime() time.Time {
	return initing
}

func RunningTime() time.Time {
	return running
}

func CostTime() time.Duration {
	return running.Sub(startup)
}
