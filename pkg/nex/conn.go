package nex

import (
	"context"
	"google.golang.org/grpc/metadata"
	"time"

	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	health "google.golang.org/grpc/health/grpc_health_v1"

	_ "github.com/mbobakov/grpc-consul-resolver"
)

type Conn struct {
	Addr string
	Info *proto.ServiceInfo

	commandServer   *grpc.Server
	commandHandlers map[string]CommandHandler

	nexusConn  *grpc.ClientConn
	clientConn map[string]*grpc.ClientConn
}

func NewNexusConn(addr string, info *proto.ServiceInfo) (*Conn, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Conn{
		Addr: addr,
		Info: info,

		commandHandlers: make(map[string]CommandHandler),

		nexusConn:  conn,
		clientConn: make(map[string]*grpc.ClientConn),
	}, nil
}

func (v *Conn) RegisterService() error {
	dir := proto.NewDirectoryServiceClient(v.nexusConn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "client_id", v.Info.Id)
	_, err := dir.AddService(ctx, v.Info)
	return err
}

func (v *Conn) RunRegistering() error {
	err := v.RegisterService()
	if err != nil {
		return err
	}

	for {
		time.Sleep(5 * time.Second)
		client := health.NewHealthClient(v.nexusConn)
		if _, err := client.Check(context.Background(), &health.HealthCheckRequest{}); err != nil {
			if v.RunRegistering() == nil {
				break
			}
		}
	}

	return nil
}

func (v *Conn) GetNexusGrpcConn() *grpc.ClientConn {
	return v.nexusConn
}

func (v *Conn) GetClientGrpcConn(t string) (*grpc.ClientConn, error) {
	if val, ok := v.clientConn[t]; ok {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if _, err := health.NewHealthClient(val).Check(ctx, &health.HealthCheckRequest{
			Service: t,
		}); err == nil {
			return val, nil
		} else {
			delete(v.clientConn, t)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	out, err := proto.NewDirectoryServiceClient(v.nexusConn).GetService(ctx, &proto.GetServiceRequest{
		Type: &t,
	})
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(
		out.GetData().GetGrpcAddr(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err == nil {
		v.clientConn[t] = conn
	}
	return conn, err
}
