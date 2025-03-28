package cache

import "github.com/redis/go-redis/v9"

var Rdb *redis.Client

func ConnectRedis(addr, password string, db int) error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return nil
}
