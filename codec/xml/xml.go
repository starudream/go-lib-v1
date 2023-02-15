package xml

import (
	"encoding/xml"
)

var (
	Prefix = ""
	Indent = "  "

	Marshal   = xml.Marshal
	Unmarshal = xml.Unmarshal

	MarshalIndent = xml.MarshalIndent

	NewDecoder = xml.NewDecoder
	NewEncoder = xml.NewEncoder
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

func MustMarshalIndent(v any) []byte {
	bs, err := MarshalIndent(v, Prefix, Indent)
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
