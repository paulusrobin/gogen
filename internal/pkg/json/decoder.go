package json

import (
	jsoniter "github.com/json-iterator/go"
	"io"
)

// NewDecoder function to return jsoniter.Decoder.
func NewDecoder(reader io.Reader) *jsoniter.Decoder {
	return json.NewDecoder(reader)
}

// Unmarshal function using jsoniter.
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// UnmarshalFromString function using jsoniter.
func UnmarshalFromString(data string, v interface{}) error {
	return json.UnmarshalFromString(data, v)
}
