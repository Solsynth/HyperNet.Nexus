package cruda_test

import (
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/cruda"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"testing"
	"time"
)

type Test struct {
	cruda.BaseModel
	Content string `json:"content" validate:"required"`
}

func TestCrudaCommand(t *testing.T) {
	conn, err := nex.NewNexusConn("127.0.0.1:7001", &proto.ServiceInfo{
		Id:       "cruda01",
		Type:     "cruda",
		Label:    "CRUD Accelerator",
		GrpcAddr: "127.0.0.1:6001",
		HttpAddr: nil,
	})
	if err != nil {
		t.Fatal(fmt.Errorf("unable to connect nexus: %v", err))
	}

	if err := conn.RegisterService(); err != nil {
		t.Fatal(fmt.Errorf("unable to register service: %v", err))
	}

	cc := cruda.NewCrudaConn(conn)
	dsn, err := cc.AllocDatabase("test")
	if err != nil {
		t.Fatal(fmt.Errorf("unable to allocate database: %v", err))
	}
	t.Log(fmt.Sprintf("Allocated database: %s", dsn))

	if err := cruda.MigrateModel(cc, Test{}); err != nil {
		t.Fatal(fmt.Errorf("unable to migrate database: %v", err))
	}

	if err := cruda.AddModel(cc, Test{}, "tm", "test.", nil); err != nil {
		t.Fatal(fmt.Errorf("unable to add commands: %v", err))
	}

	go func() {
		err := conn.RunCommands("0.0.0.0:6001")
		if err != nil {
			t.Error(fmt.Errorf("unable to run commands: %v", err))
			return
		}
	}()

	t.Log("Waiting 180 seconds for calling command...")
	time.Sleep(time.Second * 180)
}
