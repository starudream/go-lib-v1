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

func UUIDShort() string {
	src, dst := []byte(UUID()), make([]byte, 32)
	copy(dst[0:8], src[0:8])
	copy(dst[8:12], src[9:13])
	copy(dst[12:16], src[14:18])
	copy(dst[16:20], src[19:23])
	copy(dst[20:32], src[24:36])
	return string(dst)
}
