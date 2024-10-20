package directory

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"time"
)

type CommandRpcServer struct {
	proto.UnimplementedCommandControllerServer
}

func (c CommandRpcServer) AddCommand(ctx context.Context, info *proto.CommandInfo) (*proto.AddCommandResponse, error) {
	clientId, err := GetClientId(ctx)
	if err != nil {
		return nil, err
	}

	service := GetServiceInstance(clientId)
	if service == nil {
		return nil, status.Errorf(codes.NotFound, "service not found")
	}

	AddCommand(info.GetId(), info.GetMethod(), info.GetTags(), service)
	return &proto.AddCommandResponse{
		IsSuccess: true,
	}, nil
}

func (c CommandRpcServer) RemoveCommand(ctx context.Context, request *proto.CommandLookupRequest) (*proto.RemoveCommandResponse, error) {
	RemoveCommand(request.GetId(), request.GetMethod())
	return &proto.RemoveCommandResponse{
		IsSuccess: true,
	}, nil
}

func (c CommandRpcServer) SendCommand(ctx context.Context, argument *proto.CommandArgument) (*proto.CommandReturn, error) {
	id := argument.GetCommand()
	method := argument.GetMethod()

	handler := GetCommandHandler(id, method)
	if handler == nil {
		return &proto.CommandReturn{
			IsDelivered: false,
			Status:      http.StatusNotFound,
			Payload:     []byte("command not found"),
		}, nil
	}

	conn, err := handler.GetGrpcConn()
	if err != nil {
		return &proto.CommandReturn{
			IsDelivered: false,
			Status:      http.StatusServiceUnavailable,
			Payload:     []byte("service unavailable"),
		}, nil
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	out, err := proto.NewCommandControllerClient(conn).SendCommand(ctx, argument)
	if err != nil {
		return &proto.CommandReturn{
			IsDelivered: true,
			Status:      http.StatusInternalServerError,
			Payload:     []byte(err.Error()),
		}, nil
	}
	out.IsDelivered = true
	return out, nil
}

func (c CommandRpcServer) SendStreamCommand(g grpc.BidiStreamingServer[proto.CommandArgument, proto.CommandReturn]) error {
	for {
		pck, err := g.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		id := pck.GetCommand()
		method := pck.GetMethod()

		handler := GetCommandHandler(id, method)
		if handler == nil {
			return status.Errorf(codes.NotFound, "command not found")
		}

		conn, err := handler.GetGrpcConn()

		ctx, cancel := context.WithTimeout(g.Context(), time.Second*10)
		result, err := proto.NewCommandControllerClient(conn).SendCommand(ctx, pck)
		cancel()

		if err != nil {
			_ = g.Send(&proto.CommandReturn{
				IsDelivered: false,
				Status:      http.StatusInternalServerError,
				Payload:     []byte(err.Error()),
			})
		} else {
			_ = g.Send(&proto.CommandReturn{
				Status:  result.Status,
				Payload: result.Payload,
			})
		}
	}
}
