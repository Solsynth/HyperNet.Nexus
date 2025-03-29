package cache

import (
	"time"

	"git.solsynth.dev/hypernet/nexus/pkg/nex/cachekit"
	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Kcc *cachekit.Conn
)

func ConnectRedis(addr, password string, db int) error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	Kcc = &cachekit.Conn{
		Rd:      Rdb,
		Timeout: 3 * time.Second,
	}
	return nil
}
