package utils

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func init() {
	// Lambda-optimized Redis settings
	Rdb = redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_HOST"),
		Password:     os.Getenv("REDIS_PASSWORD"),
		DB:           0,

		// Lambda friendly settings:
		// Lambda freezes processes, so idle connections MUST be low.
		PoolSize:     3,              // keep small
		MinIdleConns: 0,              // Lambda doesn't keep warm idle conns
		MaxIdleConns: 2,              // avoid "idle connections timed out"

		DialTimeout:  2 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	})
}

func Set(key string, value interface{}, ttl time.Duration) error {
	return Rdb.Set(Ctx, key, value, ttl).Err()
}

func Get(key string) (string, error) {
	return Rdb.Get(Ctx, key).Result()
}
