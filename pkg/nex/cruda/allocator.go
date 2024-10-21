package cruda

import (
	"context"
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"google.golang.org/grpc/metadata"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (v *CrudConn) AllocDatabase(name string) (string, error) {
	conn := v.Conn.GetNexusGrpcConn()
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "client_id", v.Conn.Info.Id)
	out, err := proto.NewDatabaseServiceClient(conn).AllocDatabase(ctx, &proto.AllocDatabaseRequest{
		Name: name,
	})
	if err != nil || !out.GetIsSuccess() {
		return "", err
	}
	dsn := out.GetDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return "", err
	}
	v.db = db
	return dsn, nil
}

func MigrateModel[T any](v *CrudConn, model T) error {
	if v.db == nil {
		return fmt.Errorf("database has not been allocated")
	}
	return v.db.AutoMigrate(model)
}
