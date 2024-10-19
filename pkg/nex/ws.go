package nex

import "github.com/goccy/go-json"

type WebSocketPackage struct {
	Action   string `json:"w"`
	Endpoint string `json:"e,omitempty"`
	Message  string `json:"m,omitempty"`
	Payload  any    `json:"p"`
}

func (v WebSocketPackage) Marshal() []byte {
	data, _ := json.Marshal(v)
	return data
}

func (v WebSocketPackage) RawPayload() []byte {
	out, _ := json.Marshal(v.Payload)
	return out
}
