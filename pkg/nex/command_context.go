package nex

import (
	"github.com/goccy/go-json"
	"sync"
)

type CommandCtx struct {
	requestBody  []byte
	responseBody []byte

	statusCode int

	values sync.Map
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

func (c *CommandCtx) Write(data []byte, statusCode ...int) error {
	c.responseBody = data
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
	return c.Write(raw, statusCode...)
}
