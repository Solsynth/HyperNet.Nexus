package api

import (
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"strings"
)

type meetRoomArgs struct {
	RoomName string       `json:"room_name"`
	User     meetRoomUser `json:"user"`
}

type meetRoomUser struct {
	Avatar string `json:"avatar"`
	Nick   string `json:"nick"`
}

func renderMeetRoom(c *fiber.Ctx) error {
	if err := sec.EnsureAuthenticated(c); err != nil {
		return err
	}
	user := c.Locals("nex_user").(*sec.UserInfo)

	channel := c.Params("channel")

	var nick string
	if val, ok := user.Metadata["nick"].(string); ok {
		nick = val
	} else {
		nick = user.Name
	}
	var avatar string
	if val, ok := user.Metadata["avatar"].(string); ok {
		if strings.HasPrefix(val, "http") {
			avatar = val
		} else {
			endpoint := viper.GetString("resources_endpoint")
			avatar = fmt.Sprintf("%s/attachments/%s", endpoint, val)
		}
		avatar = fmt.Sprintf("\"%s\"", avatar) // Make the avatar a string to embed into the js
	} else {
		avatar = "undefined"
	}

	return c.Render("meet", meetRoomArgs{
		RoomName: fmt.Sprintf("%s-%s", "sn-chat", channel),
		User: meetRoomUser{
			Avatar: avatar,
			Nick:   nick,
		},
	})
}
