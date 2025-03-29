package cachekit

import (
	"context"
	"fmt"
	"sync"
	"time"

	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"github.com/redis/go-redis/v9"
)

// The global variable below is used to keep there will only be one redis client exist in a single instance
// Prevent if other DirectAccess™ SDK creating too many redis clients
// And able to recreate the conn with different options
var (
	rdc *redis.Client
	rdl *sync.Mutex
)

type Conn struct {
	n       *nex.Conn
	Rd      *redis.Client
	Timeout time.Duration
}

func NewConn(conn *nex.Conn, timeout time.Duration) (*Conn, error) {
	rdl.Lock()
	defer rdl.Unlock()

	c := &Conn{
		n:       conn,
		Timeout: timeout,
	}

	if rdc != nil {
		c.Rd = rdc
		return c, nil
	}

	rdb := conn.AllocResource(nex.AllocatableResourceCache)
	if rdb == nil {
		return nil, fmt.Errorf("unable to allocate resource: cache")
	} else if client, ok := rdb.(*redis.Client); !ok {
		return nil, fmt.Errorf("allocated cache resource is not a redis client")
	} else {
		c.Rd = client
		rdc = client
	}

	return c, nil
}

func (c *Conn) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), c.Timeout)
}
