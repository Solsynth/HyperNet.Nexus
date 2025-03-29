package cachekit

import (
	redis_store "github.com/eko/gocache/store/redis/v4"
)

func (c *Conn) GoCache() *redis_store.RedisStore {
	return redis_store.NewRedis(c.Rd)
}
