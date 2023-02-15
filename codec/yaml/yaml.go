package yaml

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

var (
	Space = 2

	Marshal   = yaml.Marshal
	Unmarshal = yaml.Unmarshal

	NewEncoder = yaml.NewEncoder
	NewDecoder = yaml.NewDecoder
)

func MarshalIndent(v any, space int) ([]byte, error) {
	w := &bytes.Buffer{}
	e := NewEncoder(w)
	defer func() { _ = e.Close() }()
	e.SetIndent(space)
	err := e.Encode(v)
	return w.Bytes(), err
}

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

func MustMarshalIndent(v any) []byte {
	bs, err := MarshalIndent(v, Space)
	if err != nil {
		panic(err)
	}
	return bs
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
