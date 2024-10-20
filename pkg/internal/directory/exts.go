package directory

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GetClientId(ctx context.Context) (string, error) {
	var clientId string
	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		return clientId, status.Errorf(codes.InvalidArgument, "missing metadata")
	} else if val, ok := md["client_id"]; !ok || len(val) == 0 {
		return clientId, status.Errorf(codes.Unauthenticated, "missing client_id in metadata")
	} else {
		clientId = val[0]
	}
	return clientId, nil
}
