package nex

import (
	"context"
	"time"

	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	health "google.golang.org/grpc/health/grpc_health_v1"

	_ "github.com/mbobakov/grpc-consul-resolver"
)

type HyperConn struct {
	Addr string
	Info *proto.ServiceInfo

	dealerConn    *grpc.ClientConn
	cacheGrpcConn map[string]*grpc.ClientConn
}

func NewHyperConn(addr string, info *proto.ServiceInfo) (*HyperConn, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &HyperConn{
		Addr: addr,
		Info: info,

		dealerConn:    conn,
		cacheGrpcConn: make(map[string]*grpc.ClientConn),
	}, nil
}

func (v *HyperConn) RegisterService() error {
	dir := proto.NewServiceDirectoryClient(v.dealerConn)
	_, err := dir.AddService(context.Background(), v.Info)
	return err
}

func (v *HyperConn) KeepRegisterService() error {
	err := v.RegisterService()
	if err != nil {
		return err
	}

	for {
		time.Sleep(5 * time.Second)
		client := health.NewHealthClient(v.dealerConn)
		if _, err := client.Check(context.Background(), &health.HealthCheckRequest{}); err != nil {
			if v.KeepRegisterService() == nil {
				break
			}
		}
	}

	return nil
}

func (v *HyperConn) GetNexusGrpcConn() *grpc.ClientConn {
	return v.dealerConn
}

func (v *HyperConn) GetServiceGrpcConn(t string) (*grpc.ClientConn, error) {
	if val, ok := v.cacheGrpcConn[t]; ok {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if _, err := health.NewHealthClient(val).Check(ctx, &health.HealthCheckRequest{
			Service: t,
		}); err == nil {
			return val, nil
		} else {
			delete(v.cacheGrpcConn, t)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	out, err := proto.NewServiceDirectoryClient(v.dealerConn).GetService(ctx, &proto.GetServiceRequest{
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
		v.cacheGrpcConn[t] = conn
	}
	return conn, err
}
