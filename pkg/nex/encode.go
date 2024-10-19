package nex

import jsoniter "github.com/json-iterator/go"

func EncodeMap(policy map[string]any) []byte {
	raw, _ := jsoniter.Marshal(policy)
	return raw
}

func DecodeMap(raw []byte) map[string]any {
	var out map[string]any
	_ = jsoniter.Unmarshal(raw, &out)
	return out
}
