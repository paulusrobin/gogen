package json

import (
	jsoniter "github.com/json-iterator/go"
	"io"
)

// NewEncoder function to return jsoniter.Encoder.
func NewEncoder(writer io.Writer) *jsoniter.Encoder {
	return json.NewEncoder(writer)
}

// Marshal function using jsoniter.
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalToString function using jsoniter.
func MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}

// MarshalIndent function using jsoniter.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}
