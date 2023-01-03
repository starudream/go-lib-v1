package json

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Codec implements the encoding.Encoder and encoding.Decoder interfaces for JSON encoding.
type Codec struct{}

func (Codec) Encode(v map[string]interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func (Codec) Decode(b []byte, v map[string]interface{}) error {
	return json.Unmarshal(b, &v)
}
