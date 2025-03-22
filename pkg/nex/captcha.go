package nex

import (
	"context"

	"git.solsynth.dev/hypernet/nexus/pkg/proto"
)

func (v *Conn) ValidateCaptcha(token, ip string) bool {
	client := proto.NewCaptchaServiceClient(v.nexusConn)
	resp, err := client.CheckCaptcha(context.Background(), &proto.CheckCaptchaRequest{
		Token:    token,
		RemoteIp: ip,
	})
	if err != nil {
		return false
	}
	return resp.GetIsValid()
}
