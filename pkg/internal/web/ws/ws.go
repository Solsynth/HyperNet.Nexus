package ws

import (
	"context"
	"fmt"

	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/gofiber/contrib/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/samber/lo"
)

func Listen(c *websocket.Conn) {
	user, ok := c.Locals("nex_user").(*sec.UserInfo)
	if !ok {
		_ = c.WriteMessage(1, nex.WebSocketPackage{
			Action:  "error",
			Message: "unauthorized",
		}.Marshal())
		c.Close()
		return
	}

	// Push connection
	var err error
	clientId, err := ClientRegister(*user, c)
	if err != nil {
		_ = c.WriteMessage(1, nex.WebSocketPackage{
			Action:  "error",
			Message: "client with this id already connected",
		}.Marshal())
		c.Close()
		return
	}

	// Event loop
	var mt int
	var data []byte
	var packet nex.WebSocketPackage

	for {
		if mt, data, err = c.ReadMessage(); err != nil {
			break
		} else if err := jsoniter.Unmarshal(data, &packet); err != nil {
			_ = c.WriteMessage(mt, nex.WebSocketPackage{
				Action:  "error",
				Message: "unable to unmarshal your command, requires json request",
			}.Marshal())
			continue
		}

		service := directory.GetServiceInstanceByType(packet.Endpoint)
		if service == nil {
			_ = c.WriteMessage(mt, nex.WebSocketPackage{
				Action:  "error",
				Message: "service not found",
			}.Marshal())
			continue
		}
		pc, err := service.GetGrpcConn()
		if err != nil {
			_ = c.WriteMessage(mt, nex.WebSocketPackage{
				Action:  "error",
				Message: fmt.Sprintf("unable to connect to service: %v", err.Error()),
			}.Marshal())
			continue
		}

		sc := proto.NewStreamServiceClient(pc)
		_, err = sc.PushStream(context.Background(), &proto.PushStreamRequest{
			UserId:   lo.ToPtr(uint64(user.ID)),
			ClientId: lo.ToPtr(clientId),
			Body:     packet.Marshal(),
		})
		if err != nil {
			_ = c.WriteMessage(mt, nex.WebSocketPackage{
				Action:  "error",
				Message: fmt.Sprintf("unable send message to service: %v", err.Error()),
			}.Marshal())
			continue
		}
	}

	// Pop connection
	ClientUnregister(*user, clientId)
}
