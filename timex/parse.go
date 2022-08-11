package timex

import (
	"time"

	"github.com/spf13/cast"
)

func Parse(v any, l ...*time.Location) (time.Time, error) {
	if len(l) == 0 {
		l = []*time.Location{time.Local}
	}
	return cast.ToTimeInDefaultLocationE(v, l[0])
}

func MustParse(v any, l ...*time.Location) time.Time {
	t, err := Parse(v, l...)
	if err != nil {
		panic(err)
	}
	return t
}
