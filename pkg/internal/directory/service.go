package directory

import (
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type ServiceInstance struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Label    string  `json:"label"`
	GrpcAddr string  `json:"grpc_addr"`
	HttpAddr *string `json:"http_addr"`

	retryCount int
}

var connectionCache = make(map[string]*grpc.ClientConn)

func (v *ServiceInstance) GetGrpcConn() (*grpc.ClientConn, error) {
	if conn, ok := connectionCache[v.ID]; ok {
		return conn, nil
	}

	conn, err := ConnectService(v)
	if err != nil {
		_ = RemoveServiceInstance(v.ID)
		log.Error().Str("id", v.ID).Err(err).Msg("Failed to connect to service, dropped...")
		return nil, err
	} else {
		connectionCache[v.ID] = conn
	}

	return conn, nil
}
