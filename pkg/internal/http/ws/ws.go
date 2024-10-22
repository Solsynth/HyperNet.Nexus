package ws

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"github.com/gofiber/contrib/websocket"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Listen(c *websocket.Conn) {
	user := c.Locals("nex_user").(sec.UserInfo)

	// Push connection
	clientId := ClientRegister(user, c)
	log.Debug().
		Uint("user", user.ID).
		Uint64("clientId", clientId).
		Msg("New websocket connection established...")

	// Event loop
	var mt int
	var data []byte
	var err error

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

		aliasingMap := viper.GetStringMapString("services.aliases")
		if val, ok := aliasingMap[packet.Endpoint]; ok {
			packet.Endpoint = val
		}

		/*
			service := directory.GetServiceInstanceByType(packet.Endpoint)
			if service == nil {
				_ = c.WriteMessage(mt, nex.NetworkPackage{
					Action:  "error",
					Message: "service not found",
				}.Marshal())
				continue
			}
			pc, err := service.GetGrpcConn()
			if err != nil {
				_ = c.WriteMessage(mt, nex.NetworkPackage{
					Action:  "error",
					Message: fmt.Sprintf("unable to connect to service: %v", err.Error()),
				}.Marshal())
				continue
			}

			sc := proto.NewStreamServiceClient(pc)
			_, err = sc.EmitStreamEvent(context.Background(), &proto.StreamEventRequest{
				Event:    packet.Action,
				UserId:   uint64(user.ID),
				ClientId: uint64(clientId),
				Payload:  packet.RawPayload(),
			})
			if err != nil {
				_ = c.WriteMessage(mt, nex.NetworkPackage{
					Action:  "error",
					Message: fmt.Sprintf("unable send message to service: %v", err.Error()),
				}.Marshal())
				continue
			}*/
	}

	// Pop connection
	ClientUnregister(user, clientId)
	log.Debug().
		Uint("user", user.ID).
		Uint64("clientId", clientId).
		Msg("A websocket connection disconnected...")
}
