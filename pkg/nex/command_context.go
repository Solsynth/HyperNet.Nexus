package nex

import (
	"github.com/goccy/go-json"
	"sync"
)

type CommandCtx struct {
	requestBody  []byte
	responseBody []byte

	contentType string
	statusCode  int

	values sync.Map
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
