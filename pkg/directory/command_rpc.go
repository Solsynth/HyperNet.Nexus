package directory

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
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

	service := GetServiceInstanceByType(clientId)
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
		return nil, status.Errorf(codes.NotFound, "command not found")
	}

	conn, err := handler.GetGrpcConn()
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "service unavailable")
	}

	contx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return proto.NewCommandControllerClient(conn).SendCommand(contx, argument)
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

		contx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		result, _ := proto.NewCommandControllerClient(conn).SendCommand(contx, pck)
		cancel()

		_ = g.Send(&proto.CommandReturn{
			Status:  result.Status,
			Payload: result.Payload,
		})
	}
}
