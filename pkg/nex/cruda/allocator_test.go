package cruda_test

import (
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/cruda"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"testing"
)

func TestAllocDatabase(t *testing.T) {
	conn, err := nex.NewNexusConn("127.0.0.1:7001", &proto.ServiceInfo{
		Id:       "alloc01",
		Type:     "alloc",
		Label:    "Allocator",
		GrpcAddr: "127.0.0.1:6001",
		HttpAddr: nil,
	})
	if err != nil {
		t.Fatal(fmt.Errorf("unable to connect nexus: %v", err))
	}

	if err := conn.RegisterService(); err != nil {
		t.Fatal(fmt.Errorf("unable to register service: %v", err))
	}

	cc := cruda.NewCudaConn(conn)
	dsn, err := cc.AllocDatabase("test")
	if err != nil {
		t.Fatal(fmt.Errorf("unable to allocate database: %v", err))
	}
	t.Log(fmt.Sprintf("Allocated database: %s", dsn))
}
