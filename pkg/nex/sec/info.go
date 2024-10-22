package sec

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/goccy/go-json"
)

// UserInfo is the basic of userinfo, you can add anything above it.
// Full data from id service was stored in the metadata field.
type UserInfo struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	PermNodes map[string]any `json:"perm_nodes"`
	Metadata  map[string]any `json:"metadata"`
}

func NewUserInfoFromProto(in *proto.UserInfo) UserInfo {
	return UserInfo{
		ID:        uint(in.Id),
		Name:      in.Name,
		PermNodes: nex.DecodeMap(in.PermNodes),
		Metadata:  nex.DecodeMap(in.Metadata),
	}
}

func NewUserInfoFromBytes(in []byte) (UserInfo, error) {
	var info UserInfo
	err := json.Unmarshal(in, &info)
	return info, err
}

func (v UserInfo) Encode() []byte {
	return nex.EncodeMap(v)
}
