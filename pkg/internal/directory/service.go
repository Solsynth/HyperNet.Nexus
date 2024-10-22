package directory

import "google.golang.org/grpc"

type ServiceInstance struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Label    string  `json:"label"`
	GrpcAddr string  `json:"grpc_addr"`
	HttpAddr *string `json:"http_addr"`

	grpcConn   *grpc.ClientConn
	retryCount int
}

func (v *ServiceInstance) GetGrpcConn() (*grpc.ClientConn, error) {
	if v.grpcConn != nil {
		return v.grpcConn, nil
	}

	var err error
	v.grpcConn, err = ConnectService(v)
	if err != nil {
		_ = RemoveServiceInstance(v.ID)
		return nil, err
	}

	return v.grpcConn, nil
}
