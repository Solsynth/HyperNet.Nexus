package grpc

import (
	"context"
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/http/ws"

	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/samber/lo"
)

func (v *Server) CountStreamConnection(ctx context.Context, request *proto.CountConnectionRequest) (*proto.CountConnectionResponse, error) {
	out := ws.ClientCount(uint(request.GetUserId()))
	return &proto.CountConnectionResponse{
		Count: int64(out),
	}, nil
}

func (v *Server) PushStream(ctx context.Context, request *proto.PushStreamRequest) (*proto.PushStreamResponse, error) {
	var cnt int
	var success int
	var errs []error
	if request.UserId != nil {
		cnt, success, errs = ws.WebsocketPush(uint(request.GetUserId()), request.GetBody())
	} else if request.ClientId != nil {
		cnt, success, errs = ws.WebsocketPushDirect(request.GetClientId(), request.GetBody())
	} else {
		return nil, fmt.Errorf("you must give one of the user id or client id")
	}

	if len(errs) > 0 {
		// Partial fail
		return &proto.PushStreamResponse{
			IsAllSuccess:  false,
			AffectedCount: int64(success),
			FailedCount:   int64(cnt - success),
		}, nil
	} else if cnt > 0 && success == 0 {
		// All fail
		return nil, fmt.Errorf("all push request failed: %v", errs)
	}

	return &proto.PushStreamResponse{
		IsAllSuccess:  true,
		AffectedCount: int64(success),
		FailedCount:   int64(cnt - success),
	}, nil
}

func (v *Server) PushStreamBatch(ctx context.Context, request *proto.PushStreamBatchRequest) (*proto.PushStreamResponse, error) {
	var cnt int
	var success int
	var errs []error
	if len(request.UserId) != 0 {
		cnt, success, errs = ws.WebsocketPushBatch(
			lo.Map(request.GetUserId(), func(item uint64, idx int) uint {
				return uint(item)
			},
			), request.GetBody(),
		)
	}
	if len(request.ClientId) != 0 {
		cCnt, cSuccess, cErrs := ws.WebsocketPushBatchDirect(request.GetClientId(), request.GetBody())
		cnt += cCnt
		success += cSuccess
		errs = append(errs, cErrs...)
	}

	if len(errs) > 0 {
		// Partial fail
		return &proto.PushStreamResponse{
			IsAllSuccess:  false,
			AffectedCount: int64(success),
			FailedCount:   int64(cnt - success),
		}, nil
	} else if cnt > 0 && success == 0 {
		// All fail
		return nil, fmt.Errorf("all push request failed: %v", errs)
	}

	return &proto.PushStreamResponse{
		IsAllSuccess:  true,
		AffectedCount: int64(success),
		FailedCount:   int64(cnt - success),
	}, nil
}
