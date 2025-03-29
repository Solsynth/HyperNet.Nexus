package cachekit

import (
	"fmt"
	"time"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/goccy/go-json"
)

// The functions below are operating the redis via the gocache
// Provide a advanced tagging experience
// At the same time, the advanced cache using client side marshaling to handle the advance data types

func Set[T any](c *CaConn, key string, value T, ttl time.Duration, tags ...string) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("unable to marshal value during caching: %v", err)
	}

	ctx, cancel := c.withTimeout()
	defer cancel()
	cm := cache.New[[]byte](c.GoCache())
	return cm.Set(ctx, key, raw, store.WithTags(tags), store.WithExpiration(ttl))
}

// SetKA stands for Set Keep Alive
// Don't set a TTL for the value set via this function
func SetKA[T any](c *CaConn, key string, value T, tags ...string) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("unable to marshal value during caching: %v", err)
	}

	ctx, cancel := c.withTimeout()
	defer cancel()
	cm := cache.New[[]byte](c.GoCache())
	return cm.Set(ctx, key, raw, store.WithTags(tags))
}

func Get[T any](c *CaConn, key string) (T, error) {
	var out T

	ctx, cancel := c.withTimeout()
	defer cancel()
	cm := cache.New[[]byte](c.GoCache())
	raw, err := cm.Get(ctx, key)
	if err != nil {
		return out, err
	}

	if err := json.Unmarshal(raw, &out); err != nil {
		return out, fmt.Errorf("unable to unmarshal value during caching: %v", err)
	}

	return out, nil
}

func Delete(c *CaConn, key string) error {
	ctx, cancel := c.withTimeout()
	defer cancel()
	cm := cache.New[[]byte](c.GoCache())
	return cm.Delete(ctx, key)
}

func DeleteByTags(c *CaConn, tags ...string) error {
	if len(tags) == 0 {
		return nil
	}
	ctx, cancel := c.withTimeout()
	defer cancel()
	cm := cache.New[[]byte](c.GoCache())
	return cm.Invalidate(ctx, store.WithInvalidateTags(tags))
}
