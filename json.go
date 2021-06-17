package codec

import (
	"bytes"
	"encoding/json"
)

// StringList stringList
type StringList []string

// MarshalJSON implement json.Marshaler interface
func (sList StringList) MarshalJSON() ([]byte, error) {

	var buf bytes.Buffer
	buf.WriteString("[")
	for i, s := range sList {
		buf.WriteString(s)
		if i+1 != len(sList) {
			buf.WriteString(",")
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}

// Bytes bytes
func (sList StringList) Bytes() ([]byte, error) {
	return sList.MarshalJSON()
}

// JSON implements Codec and Formatter interface
type JSON struct{}

// Encode implate Encoder interface
func (JSON) Encode(v interface{}) ([]byte, error) {
	switch v.(type) {
	case []byte:
		return v.([]byte), nil
	case string:
		return []byte(v.(string)), nil
	case []string:
		return StringList(v.([]string)).Bytes()
	default:
		return json.Marshal(v)
	}
}

// Decode implate Decoder interface
func (JSON) Decode(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

// Format implate Formater interface
func (json JSON) Format(dest, src interface{}) error {
	b, err := json.Encode(src)
	if err != nil {
		return err
	}
	return json.Decode(b, dest)
}

func (JSON) String() string {
	return "json"
}

// NewJSONCodec jsoncodec
func NewJSONCodec() JSON {
	return JSON{}
}
