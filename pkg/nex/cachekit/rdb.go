package cachekit

import (
	"context"
	"fmt"
	"time"

	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"github.com/redis/go-redis/v9"
)

type CaConn struct {
	n       *nex.Conn
	Rd      *redis.Client
	Timeout time.Duration
}

func NewCaConn(conn *nex.Conn, timeout time.Duration) (*CaConn, error) {
	c := &CaConn{
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

func (c *CaConn) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), c.Timeout)
}

// Set stores a key-value pair in Redis with an optional expiration time
func (c *CaConn) Set(key string, value any, ttl time.Duration) error {
	ctx, cancel := c.withTimeout()
	defer cancel()
	return c.Rd.Set(ctx, key, value, ttl).Err()
}

// Get retrieves a value from Redis by key
func (c *CaConn) Get(key string) (string, error) {
	ctx, cancel := c.withTimeout()
	defer cancel()
	return c.Rd.Get(ctx, key).Result()
}

// Delete removes a key from Redis
func (c *CaConn) Delete(key string) error {
	ctx, cancel := c.withTimeout()
	defer cancel()
	return c.Rd.Del(ctx, key).Err()
}

// Exists checks if a key exists in Redis
func (c *CaConn) Exists(key string) (bool, error) {
	ctx, cancel := c.withTimeout()
	defer cancel()
	exists, err := c.Rd.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

// ClearCacheByPrefix deletes all keys matching a given prefix
func (c *CaConn) DeleteByPrefix(prefix string) error {
	ctx, cancel := c.withTimeout()
	defer cancel()

	iter := c.Rd.Scan(ctx, 0, prefix+"*", 0).Iterator()
	for iter.Next(ctx) {
		if err := c.Rd.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}
