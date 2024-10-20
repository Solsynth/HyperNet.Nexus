package nex

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"strings"
	"time"
)

type CommandHandler func(ctx *CommandCtx) error

func GetCommandKey(id, method string) string {
	return id + ":" + method
}

func (v *Conn) AddCommand(id, method string, tags []string, fn CommandHandler) error {
	method = strings.ToLower(method)
	dir := proto.NewCommandControllerClient(v.nexusConn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "client_id", v.Info.Id)

	var addingMethodQueue []string
	if method == "all" {
		addingMethodQueue = []string{"get", "post", "put", "patch", "delete"}
	} else {
		addingMethodQueue = append(addingMethodQueue, method)
	}

	for _, method := range addingMethodQueue {
		ky := GetCommandKey(id, method)
		_, err := dir.AddCommand(ctx, &proto.CommandInfo{
			Id:     id,
			Method: method,
			Tags:   tags,
		})
		if err == nil {
			v.commandHandlers[ky] = fn
		} else {
			return err
		}
	}

	return nil
}

type localCommandRpcServer struct {
	conn *Conn

	proto.UnimplementedCommandControllerServer
	health.UnimplementedHealthServer
}

func (v localCommandRpcServer) SendCommand(ctx context.Context, argument *proto.CommandArgument) (*proto.CommandReturn, error) {
	ky := GetCommandKey(argument.GetCommand(), argument.GetMethod())
	if handler, ok := v.conn.commandHandlers[ky]; !ok {
		return &proto.CommandReturn{
			Status:  http.StatusNotFound,
			Payload: []byte(argument.GetCommand() + " not found"),
		}, nil
	} else {
		cc := &CommandCtx{
			requestBody: argument.GetPayload(),
			statusCode:  http.StatusOK,
		}
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for k, v := range md {
				cc.values.Store(k, v)
			}
		}
		if err := handler(cc); err != nil {
			return nil, err
		} else {
			return &proto.CommandReturn{
				Status:  int32(cc.statusCode),
				Payload: cc.responseBody,
			}, nil
		}
	}
}

func (v localCommandRpcServer) Check(ctx context.Context, request *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (v localCommandRpcServer) Watch(request *health.HealthCheckRequest, server health.Health_WatchServer) error {
	for {
		if server.Send(&health.HealthCheckResponse{
			Status: health.HealthCheckResponse_SERVING,
		}) != nil {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (v *Conn) RunCommands(addr string) error {
	v.commandServer = grpc.NewServer()
	service := &localCommandRpcServer{conn: v}
	proto.RegisterCommandControllerServer(v.commandServer, service)
	health.RegisterHealthServer(v.commandServer, service)
	reflection.Register(v.commandServer)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return v.commandServer.Serve(listener)
}
