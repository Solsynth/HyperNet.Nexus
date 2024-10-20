package cruda

import "git.solsynth.dev/hypernet/nexus/pkg/nex"

type CudaConn struct {
	Conn *nex.Conn
}

func NewCudaConn(conn *nex.Conn) *CudaConn {
	return &CudaConn{
		Conn: conn,
	}
}
