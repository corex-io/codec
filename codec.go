package codec

// Encoder encode Type to Bytes
type Encoder interface {
	Encode(interface{}) ([]byte, error)
}

// Decoder decode Bytes to Type
type Decoder interface {
	Decode([]byte, interface{}) error
}

// Codec encode/decode interface
type Codec interface {
	Encoder
	Decoder
	Formatter
}

// Formatter format v to v
type Formatter interface {
	Format(dest interface{}, src interface{}) error
}

// New codec factory
func New(kinds ...string) Codec {
	if len(kinds) == 0 {
		return NewJSONCodec()
	}
	switch kinds[0] {
	case ".json":
		return NewJSONCodec()
	case ".yaml", ".yml":
		return NewYAMLCodec()
	default:
		return NewJSONCodec()
	}
}

var codec = NewJSONCodec()

// Encode xx
func Encode(v interface{}) ([]byte, error) {
	return codec.Encode(v)
}

// Decode implate Decoder interface
func Decode(b []byte, v interface{}) error {
	return codec.Decode(b, v)
}

// Format implate Formater interface
func Format(dest, src interface{}) error {
	return codec.Format(dest, src)
}
