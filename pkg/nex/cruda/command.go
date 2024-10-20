package cruda

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"net/http"
	"strconv"
)

type CrudAction func(v *CrudConn) nex.CommandHandler

func AddModel[T any](v *CrudConn, model T, id, prefix string, tags []string) error {
	funcList := []CrudAction{cmdList[T]}
	funcCmds := []string{".list", ".get", ".create", ".update", ".delete"}
	for idx, fn := range funcList {
		if err := v.Conn.AddCommand(prefix+id+funcCmds[idx], "get", tags, fn(v)); err != nil {
			return err
		}
	}
	return nil
}

func cmdList[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		rawTake := ctx.ValueOrElse("query.take", "10").(string)
		rawSkip := ctx.ValueOrElse("query.skip", "0").(string)
		take, err := strconv.Atoi(rawTake)
		if err != nil {
			take = 10
		}
		skip, err := strconv.Atoi(rawSkip)
		if err != nil {
			skip = 0
		}

		var out []T
		if err := c.db.Offset(skip).Limit(take).Find(&out).Error; err != nil {
			return err
		}

		return ctx.JSON(out, http.StatusOK)
	}
}
