package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"git.solsynth.dev/hypernet/nexus/pkg/internal/cache"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/cachekit"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func userinfoFetch(c *fiber.Ctx) error {
	claims, ok := c.Locals("nex_principal").(*sec.JwtClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "user principal data was not found")
	}

	if val, err := cachekit.Get[sec.UserInfo](
		cache.Kcc,
		cachekit.FKey(cachekit.DAUserInfoPrefix, claims.Session),
	); err == nil {
		c.Locals("nex_user", &val)
		tk, err := IWriter.WriteUserInfoJwt(val)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("unable to sign userinfo: %v", err))
		}
		c.Locals("nex_token", tk)
		return nil
	}

	service := directory.GetServiceInstanceByType(nex.ServiceTypeAuth)
	if service != nil {
		conn, err := service.GetGrpcConn()
		if err != nil {
			log.Warn().Str("id", service.ID).Err(err).Msg("Unable to fetch userinfo, the implementation of id service is down")
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			sed, err := strconv.Atoi(claims.Session)
			if err != nil {
				return fiber.NewError(fiber.StatusUnauthorized, "invalid token payload")
			}
			resp, err := proto.NewAuthServiceClient(conn).Authenticate(ctx, &proto.AuthRequest{
				SessionId: uint64(sed),
			})
			if err != nil {
				return fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("unable to load userinfo: %v", err))
			}
			userinfo := sec.NewUserInfoFromProto(resp.Info.Info)
			c.Locals("nex_user", &userinfo)
			tk, err := IWriter.WriteUserInfoJwt(userinfo)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("unable to sign userinfo: %v", err))
			}
			c.Locals("nex_token", tk)
		}
	} else {
		log.Warn().Msg("Unable to fetch userinfo, no implementation of id service")
	}
	return nil
}
