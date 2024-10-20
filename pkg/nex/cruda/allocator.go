package cruda

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"google.golang.org/grpc/metadata"
)

func (v *CudaConn) AllocDatabase(name string) (string, error) {
	conn := v.Conn.GetNexusGrpcConn()
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "client_id", v.Conn.Info.Id)
	out, err := proto.NewDatabaseControllerClient(conn).AllocDatabase(ctx, &proto.AllocDatabaseRequest{
		Name: name,
	})
	if err != nil || !out.GetIsSuccess() {
		return "", err
	}
	return out.GetDsn(), nil
}
