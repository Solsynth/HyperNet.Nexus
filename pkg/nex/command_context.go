package nex

import (
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
	"sync"
)

type CommandCtx struct {
	requestBody  []byte
	responseBody []byte

	contentType string
	statusCode  int

	values sync.Map
}

func CtxValueMustBe[T any](c *CommandCtx, key string) (T, error) {
	if val, ok := c.values.Load(key); ok {
		if v, ok := val.(T); ok {
			return v, nil
		}
	}
	var out T
	if err := c.Write([]byte(fmt.Sprintf("value %s not found in type %T", key, out)), "text/plain+error", http.StatusBadRequest); err != nil {
		return out, err
	}
	return out, fmt.Errorf("value %s not found", key)
}

func CtxValueShouldBe[T any](c *CommandCtx, key string, defaultValue T) T {
	if val, ok := c.values.Load(key); ok {
		if v, ok := val.(T); ok {
			return v
		}
	}
	return defaultValue
}

func (c *CommandCtx) Values() map[string]any {
	duplicate := make(map[string]any)
	c.values.Range(func(key, value any) bool {
		duplicate[key.(string)] = value
		return true
	})
	return duplicate
}

func (c *CommandCtx) ValueOrElse(key string, defaultValue any) any {
	val, _ := c.values.Load(key)
	if val == nil {
		return defaultValue
	}
	return val
}

func (c *CommandCtx) Value(key string, newValue ...any) any {
	if len(newValue) > 0 {
		c.values.Store(key, newValue[0])
	}
	val, _ := c.values.Load(key)
	return val
}

func (c *CommandCtx) Read() []byte {
	return c.requestBody
}

func (c *CommandCtx) ReadJSON(out any) error {
	return json.Unmarshal(c.requestBody, out)
}

func (c *CommandCtx) Write(data []byte, contentType string, statusCode ...int) error {
	c.responseBody = data
	c.contentType = contentType
	if len(statusCode) > 0 {
		c.statusCode = statusCode[0]
	}
	return nil
}

func (c *CommandCtx) JSON(data any, statusCode ...int) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.Write(raw, "application/json", statusCode...)
}
