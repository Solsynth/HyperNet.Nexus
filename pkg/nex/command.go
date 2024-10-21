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
	"strconv"
	"strings"
	"time"
)

type CommandHandler func(ctx *CommandCtx) error

func GetCommandKey(id, method string) string {
	return id + ":" + method
}

func (v *Conn) AddCommand(id, method string, tags []string, fn CommandHandler) error {
	method = strings.ToLower(method)
	dir := proto.NewCommandProviderClient(v.nexusConn)
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

	proto.UnimplementedCommandProviderServer
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
				var val any = nil
				if len(v) == 1 {
					if len(v[0]) != 0 {
						if i, err := strconv.ParseInt(v[0], 10, 64); err == nil {
							val = i
						} else if b, err := strconv.ParseBool(v[0]); err == nil {
							val = b
						} else if f, err := strconv.ParseFloat(v[0], 64); err == nil {
							val = f
						}
						layouts := []string{
							time.RFC3339,
							"2006-01-02 15:04:05", // Example: 2024-10-20 14:55:05
							"2006-01-02",          // Example: 2024-10-20
						}
						for _, layout := range layouts {
							if t, err := time.Parse(layout, v[0]); err == nil {
								val = t
							}
						}
						if val == nil {
							val = v[0]
						}
					} else {
						val = v[0]
					}
				} else if len(v) > 1 {
					val = v
				}
				cc.values.Store(k, val)
			}
		}
		if err := handler(cc); err != nil {
			return nil, err
		} else {
			return &proto.CommandReturn{
				Status:      int32(cc.statusCode),
				ContentType: cc.contentType,
				Payload:     cc.responseBody,
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
	proto.RegisterCommandProviderServer(v.commandServer, service)
	health.RegisterHealthServer(v.commandServer, service)
	reflection.Register(v.commandServer)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return v.commandServer.Serve(listener)
}
