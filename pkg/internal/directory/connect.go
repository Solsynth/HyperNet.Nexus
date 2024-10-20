package directory

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"time"
)

func ConnectService(in *ServiceInstance) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		in.GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc connection: %v", err)
	}

	client := health.NewHealthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if _, err = client.Check(ctx, &health.HealthCheckRequest{}); err != nil {
		return conn, fmt.Errorf("grpc service is down: %v", err)
	}

	return conn, nil
}
