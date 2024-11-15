package nex_test

import (
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"net/http"
	"testing"
	"time"
)

func TestHandleCommand(t *testing.T) {
	conn, err := nex.NewNexusConn("127.0.0.1:7001", &proto.ServiceInfo{
		Id:       "echo01",
		Type:     "echo",
		Label:    "Echo",
		GrpcAddr: "127.0.0.1:6001",
		HttpAddr: nil,
	})
	if err != nil {
		t.Fatal(fmt.Errorf("unable to connect nexus: %v", err))
	}

	if err := conn.RegisterService(); err != nil {
		t.Fatal(fmt.Errorf("unable to register service: %v", err))
	}

	err = conn.AddCommand("say.hi", "all", nil, func(ctx *nex.CommandCtx) error {
		return ctx.Write([]byte("Hello, World!"), "text/plain", http.StatusOK)
	})
	if err != nil {
		t.Fatal(fmt.Errorf("unable to add command: %v", err))
		return
	}
	err = conn.AddCommand("echo", "all", nil, func(ctx *nex.CommandCtx) error {
		t.Log("Received command: ", string(ctx.Read()))
		return ctx.Write(ctx.Read(), "text/plain", http.StatusOK)
	})
	err = conn.AddCommand("echo.details", "all", nil, func(ctx *nex.CommandCtx) error {
		return ctx.JSON(map[string]any{
			"values": ctx.Values(),
			"body":   ctx.Read(),
		}, http.StatusOK)
	})
	if err != nil {
		t.Fatal(fmt.Errorf("unable to add command: %v", err))
		return
	}

	go func() {
		err := conn.RunCommands("0.0.0.0:6001")
		if err != nil {
			t.Error(fmt.Errorf("unable to run commands: %v", err))
			return
		}
	}()

	t.Log("Waiting 60 seconds for calling command...")
	time.Sleep(time.Second * 60)
}
