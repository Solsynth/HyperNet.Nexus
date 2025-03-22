package grpc

import (
	"context"

	"git.solsynth.dev/hypernet/nexus/pkg/internal/captcha"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
)

func (v *Server) CheckCaptcha(_ context.Context, req *proto.CheckCaptchaRequest) (*proto.CheckCaptchaResponse, error) {
	return &proto.CheckCaptchaResponse{
		IsValid: captcha.Validate(req.Token, req.RemoteIp),
	}, nil
}
