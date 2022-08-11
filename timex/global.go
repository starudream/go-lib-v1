package timex

import (
	"time"
)

func BeginOfMinute() time.Time {
	return New().BeginOfMinute()
}

func BeginOfHour() time.Time {
	return New().BeginOfHour()
}

func BeginOfDay() time.Time {
	return New().BeginOfDay()
}

func BeginOfWeek(weekStartDays ...time.Weekday) time.Time {
	return New().BeginOfWeek(weekStartDays...)
}

func BeginOfMonth() time.Time {
	return New().BeginOfMonth()
}

func BeginOfQuarter() time.Time {
	return New().BeginOfQuarter()
}

func BeginOfHalf() time.Time {
	return New().BeginOfHalf()
}

func BeginOfYear() time.Time {
	return New().BeginOfYear()
}

func EndOfMinute() time.Time {
	return New().EndOfMinute()
}

func EndOfHour() time.Time {
	return New().EndOfHour()
}

func EndOfDay() time.Time {
	return New().EndOfDay()
}

func EndOfWeek(weekStartDays ...time.Weekday) time.Time {
	return New().EndOfWeek(weekStartDays...)
}

func EndOfMonth() time.Time {
	return New().EndOfMonth()
}

func EndOfQuarter() time.Time {
	return New().EndOfQuarter()
}

func EndOfHalf() time.Time {
	return New().EndOfHalf()
}

func EndOfYear() time.Time {
	return New().EndOfYear()
}
