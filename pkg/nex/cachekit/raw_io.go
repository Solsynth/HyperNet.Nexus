package cachekit

import "time"

// The functions below are directly using the redis connection to operaete the redis

// Set stores a key-value pair in Redis with an optional expiration time
func (c *Conn) RSet(key string, value any, ttl time.Duration) error {
	ctx, cancel := c.withTimeout()
	defer cancel()
	return c.Rd.Set(ctx, key, value, ttl).Err()
}

// Get retrieves a value from Redis by key
func (c *Conn) RGet(key string) (string, error) {
	ctx, cancel := c.withTimeout()
	defer cancel()
	return c.Rd.Get(ctx, key).Result()
}

// Delete removes a key from Redis
func (c *Conn) RDelete(key string) error {
	ctx, cancel := c.withTimeout()
	defer cancel()
	return c.Rd.Del(ctx, key).Err()
}

// Exists checks if a key exists in Redis
func (c *Conn) RExists(key string) (bool, error) {
	ctx, cancel := c.withTimeout()
	defer cancel()
	exists, err := c.Rd.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

// ClearCacheByPrefix deletes all keys matching a given prefix
func (c *Conn) RDeleteByPrefix(prefix string) error {
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
