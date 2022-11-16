package seq

import (
	"github.com/google/uuid"
)

func init() {
	uuid.EnableRandPool()
}

func UUID() string {
	s, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return s.String()
}

var ParseUUID = uuid.Parse
