package codec

import "gopkg.in/yaml.v2"

// YAML implements Codec and Formatter interface
type YAML struct{}

// Encode implate Encoder interface
func (YAML) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// Decode implate Decoder interface
func (YAML) Decode(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}

// Format implate Formater interface
func (yaml YAML) Format(dest, src interface{}) error {
	b, err := yaml.Encode(src)
	if err != nil {
		return err
	}
	return yaml.Decode(b, dest)
}

func (YAML) String() string {
	return "yaml"
}

// NewYAMLCodec yamlcodec
func NewYAMLCodec() YAML {
	return YAML{}
}
