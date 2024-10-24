package rx

import "git.solsynth.dev/hypernet/nexus/pkg/nex"

func (v *MqConn) Publish(topic string, data any) error {
	return v.Nt.Publish(topic, nex.EncodeMap(data))
}
