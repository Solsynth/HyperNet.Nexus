package cruda

import (
	"errors"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

type CrudAction func(v *CrudConn) nex.CommandHandler

func AddModel[T any](v *CrudConn, model T, id, prefix string, tags []string) error {
	funcList := []CrudAction{cmdList[T], cmdGet[T], cmdCreate[T], cmdUpdate[T], cmdDelete[T]}
	funcCmds := []string{".list", "", "", "", ""}
	funcMethods := []string{"get", "get", "put", "patch", "delete"}
	for idx, fn := range funcList {
		if err := v.n.AddCommand(prefix+id+funcCmds[idx], funcMethods[idx], tags, fn(v)); err != nil {
			return err
		}
	}
	return nil
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func cmdList[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		take := int(nex.CtxValueShouldBe[int64](ctx, "query.take", 10))
		skip := int(nex.CtxValueShouldBe[int64](ctx, "query.skip", 0))

		var str T
		var count int64
		if err := c.Db.Model(str).Count(&count).Error; err != nil {
			return err
		}

		var out []T
		if err := c.Db.Offset(skip).Limit(take).Find(&out).Error; err != nil {
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
		id, err := nex.CtxValueMustBe[int64](ctx, "query.id")
		if err != nil {
			return err
		}

		var out T
		if err := c.Db.First(&out, "id = ?", id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.Write([]byte(err.Error()), "text/plain", http.StatusNotFound)
			}
			return err
		}

		return ctx.JSON(out, http.StatusOK)
	}
}

func cmdCreate[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		var payload T
		if err := ctx.ReadJSON(&payload); err != nil {
			return err
		} else if err := validate.Struct(payload); err != nil {
			return ctx.Write([]byte(err.Error()), "text/plain+error", http.StatusBadRequest)
		}

		if err := c.Db.Create(&payload).Error; err != nil {
			return err
		}

		return ctx.JSON(payload, http.StatusOK)
	}
}

func cmdUpdate[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		id, err := nex.CtxValueMustBe[int64](ctx, "query.id")
		if err != nil {
			return err
		}

		var payload T
		if err := ctx.ReadJSON(&payload); err != nil {
			return err
		} else if err := validate.Struct(payload); err != nil {
			return ctx.Write([]byte(err.Error()), "text/plain+error", http.StatusBadRequest)
		}

		var out T
		if err := c.Db.Model(out).Where("id = ?", id).Updates(&payload).Error; err != nil {
			return err
		}

		if err := c.Db.First(&out, "id = ?", id).Error; err != nil {
			return err
		}

		return ctx.JSON(out, http.StatusOK)
	}
}

func cmdDelete[T any](c *CrudConn) nex.CommandHandler {
	return func(ctx *nex.CommandCtx) error {
		id, err := nex.CtxValueMustBe[int64](ctx, "query.id")
		if err != nil {
			return err
		}

		var out T
		if err := c.Db.Delete(&out, "id = ?", id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.Write([]byte(err.Error()), "text/plain", http.StatusNotFound)
			}
			return err
		}

		return ctx.Write(nil, "text/plain", http.StatusOK)
	}
}
