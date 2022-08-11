package timex

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	t.Log(BeginOfMinute())
	t.Log(BeginOfHour())
	t.Log(BeginOfDay())
	t.Log(BeginOfWeek())
	t.Log(BeginOfWeek(time.Sunday))
	t.Log(BeginOfMonth())
	t.Log(BeginOfQuarter())
	t.Log(BeginOfHalf())
	t.Log(BeginOfYear())

	t.Log(EndOfMinute())
	t.Log(EndOfHour())
	t.Log(EndOfDay())
	t.Log(EndOfWeek())
	t.Log(EndOfWeek(time.Sunday))
	t.Log(EndOfMonth())
	t.Log(EndOfQuarter())
	t.Log(EndOfHalf())
	t.Log(EndOfYear())
}
