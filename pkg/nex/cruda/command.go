package cruda

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"net/http"
)

type CrudAction func(v *CrudConn) nex.CommandHandler

func AddModel[T any](v *CrudConn, model T, id, prefix string, tags []string) error {
	funcList := []CrudAction{cmdList[T], cmdGet[T], cmdCreate[T], cmdUpdate[T], cmdDelete[T]}
	funcCmds := []string{".list", ".get", ".create", ".update", ".delete"}
	funcMethods := []string{"get", "get", "put", "patch", "delete"}
	for idx, fn := range funcList {
		if err := v.Conn.AddCommand(prefix+id+funcCmds[idx], funcMethods[idx], tags, fn(v)); err != nil {
			return err
		}
	}
	return nil
}

func cmdList[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		take := int(ctx.ValueOrElse("query.take", 10).(int64))
		skip := int(ctx.ValueOrElse("query.skip", 0).(int64))

		var str T
		var count int64
		if err := c.db.Model(str).Count(&count).Error; err != nil {
			return err
		}

		var out []T
		if err := c.db.Offset(skip).Limit(take).Find(&out).Error; err != nil {
			return err
		}

		return ctx.JSON(map[string]any{
			"count": count,
			"data":  out,
		}, http.StatusOK)
	}
}

func cmdGet[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		id := ctx.ValueOrElse("query.id", 0).(int64)

		var out T
		if err := c.db.First(&out, "id = ?", id).Error; err != nil {
			return err
		}

		return ctx.JSON(out, http.StatusOK)
	}
}

func cmdCreate[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		var out T
		if err := ctx.ReadJSON(&out); err != nil {
			return err
		}
		// TODO validation

		if err := c.db.Create(&out).Error; err != nil {
			return err
		}

		return ctx.JSON(out, http.StatusOK)
	}
}

func cmdUpdate[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		id := ctx.ValueOrElse("query.id", 0).(int64)

		var payload T
		if err := ctx.ReadJSON(&payload); err != nil {
			return err
		}
		// TODO validation

		var out T
		if err := c.db.Model(out).Where("id = ?", id).Updates(&payload).Error; err != nil {
			return err
		}

		if err := c.db.First(&out, "id = ?", id).Error; err != nil {
			return err
		}

		return ctx.JSON(out, http.StatusOK)
	}
}

func cmdDelete[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		id := ctx.ValueOrElse("query.id", 0).(int64)

		var out T
		if err := c.db.Delete(&out, "id = ?", id).Error; err != nil {
			return err
		}

		return ctx.JSON(out, http.StatusOK)
	}
}
