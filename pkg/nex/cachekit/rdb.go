package cachekit

import (
	"context"
	"fmt"
	"time"

	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"github.com/redis/go-redis/v9"
)

type Conn struct {
	n       *nex.Conn
	Rd      *redis.Client
	Timeout time.Duration
}

func NewCaConn(conn *nex.Conn, timeout time.Duration) (*Conn, error) {
	c := &Conn{
		n:       conn,
		Timeout: timeout,
	}

	rdb := conn.AllocResource(nex.AllocatableResourceCache)
	if rdb == nil {
		return nil, fmt.Errorf("unable to allocate resource: cache")
	} else if client, ok := rdb.(*redis.Client); !ok {
		return nil, fmt.Errorf("allocated cache resource is not a redis client")
	} else {
		c.Rd = client
	}

	return c, nil
}

func (c *Conn) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), c.Timeout)
}
