package randx

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var _f *gofakeit.Faker

func init() {
	_f = gofakeit.New(time.Now().UnixNano())
	gofakeit.SetGlobalFaker(_f)
}

func F() *gofakeit.Faker {
	return _f
}
