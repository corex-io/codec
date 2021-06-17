package codec

import (
	"testing"
)

func TestJSONFormat(t *testing.T) {
	a := map[string]interface{}{
		"a": "b",
		"b": "c",
		"c": 1,
	}
	b := map[string]interface{}{
		"b": "d",
		"c": "d",
		"d": "e",
	}
	err := codec.Format(&a, &b)
	t.Logf("%#v, err=%v", a, err) // map[string]interface {}{"a":"b", "b":"d", "c":"d", "d":"e"}
}
