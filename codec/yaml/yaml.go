package yaml

import (
	"gopkg.in/yaml.v3"
)

var (
	Marshal   = yaml.Marshal
	Unmarshal = yaml.Unmarshal

	NewEncoder = yaml.NewEncoder
	NewDecoder = yaml.NewDecoder
)

func MustMarshal(v any) []byte {
	bs, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return bs
}

func MustMarshalString(v any) string {
	return string(MustMarshal(v))
}

func UnmarshalTo[T any](bs []byte) (m T, err error) {
	return m, Unmarshal(bs, &m)
}

func MustUnmarshalTo[T any](bs []byte) (m T) {
	err := Unmarshal(bs, &m)
	if err != nil {
		panic(err)
	}
	return m
}
