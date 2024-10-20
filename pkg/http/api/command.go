package api

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/directory"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"
)

func invokeCommand(c *fiber.Ctx) error {
	command := c.Params("command")
	method := strings.ToLower(c.Method())

	handler := directory.GetCommandHandler(command, method)
	if handler == nil {
		return fiber.NewError(fiber.StatusNotFound, "command not found")
	}

	conn, err := handler.GetGrpcConn()
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "service unavailable")
	}

	log.Debug().Str("id", command).Str("method", method).Msg("Invoking command from HTTP Gateway...")

	ctx := metadata.AppendToOutgoingContext(c.Context(), "client_id", "http-gateway", "ip", c.IP(), "user_agent", c.Get(fiber.HeaderUserAgent))
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	out, err := proto.NewCommandControllerClient(conn).SendCommand(ctx, &proto.CommandArgument{
		Command: command,
		Method:  method,
		Payload: c.Body(),
	})

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	} else {
		if !out.IsDelivered {
			log.Debug().Str("id", command).Str("method", method).Msg("Invoking command from HTTP Gateway... failed, delivery not confirmed")
		}
		return c.Status(int(out.Status)).Send(out.Payload)
	}
}
