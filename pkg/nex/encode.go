package nex

import (
	"github.com/goccy/go-json"
)

func EncodeMap(data any) []byte {
	raw, _ := json.Marshal(data)
	return raw
}

func DecodeMap(raw []byte) map[string]any {
	var out map[string]any
	_ = json.Unmarshal(raw, &out)
	return out
}
