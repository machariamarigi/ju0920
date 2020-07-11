package data

import (
	"encoding/json"
	"io"
)

// ToJSON serializes contents of the collection to JSON using json package's NewEncoder
// https://golang.org/pkg/encoding/json/#NewEncoder
func ToJSON(intf interface{}, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(intf)
}

// FromJSON decodes json serialized content using json package's NeWDecoder
// https://golang.org/pkg/encoding/json/#NewDecoder
func FromJSON(intf interface{}, reader io.Reader) error{
	decoder := json.NewDecoder(reader)
	return decoder.Decode(intf)
}
