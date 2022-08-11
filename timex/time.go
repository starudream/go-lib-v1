package timex

import (
	"time"
)

type Time struct {
	time.Time
}

func New(ts ...time.Time) *Time {
	if len(ts) == 0 {
		return &Time{Time: time.Now()}
	}
	return &Time{Time: ts[0]}
}

func (t *Time) BeginOfMinute() time.Time {
	return t.Truncate(time.Minute)
}

func (t *Time) BeginOfHour() time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Time.Hour(), 0, 0, 0, t.Time.Location())
}

func (t *Time) BeginOfDay() time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Time.Location())
}

func (t *Time) BeginOfWeek(weekStartDays ...time.Weekday) time.Time {
	weekStartDay := time.Monday
	if len(weekStartDays) > 0 {
		weekStartDay = weekStartDays[0]
	}

	_t := t.BeginOfDay()
	weekday := int(_t.Weekday())

	if weekStartDay != time.Sunday {
		weekStartDayInt := int(weekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}

	return _t.AddDate(0, 0, -weekday)
}

func (t *Time) BeginOfMonth() time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

func (t *Time) BeginOfQuarter() time.Time {
	month := t.BeginOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

func (t *Time) BeginOfHalf() time.Time {
	month := t.BeginOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

func (t *Time) BeginOfYear() time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

func (t *Time) EndOfMinute() time.Time {
	return t.BeginOfMinute().Add(time.Minute - time.Nanosecond)
}

func (t *Time) EndOfHour() time.Time {
	return t.BeginOfHour().Add(time.Hour - time.Nanosecond)
}

func (t *Time) EndOfDay() time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

func (t *Time) EndOfWeek(weekStartDays ...time.Weekday) time.Time {
	return t.BeginOfWeek(weekStartDays...).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func (t *Time) EndOfMonth() time.Time {
	return t.BeginOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (t *Time) EndOfQuarter() time.Time {
	return t.BeginOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

func (t *Time) EndOfHalf() time.Time {
	return t.BeginOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
}

func (t *Time) EndOfYear() time.Time {
	return t.BeginOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}
