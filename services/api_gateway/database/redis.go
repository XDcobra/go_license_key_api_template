package database

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectionRedisDB() *redis.Client {
	var rdb *redis.Client
	if os.Getenv("DEV") != "true" {
		rdb = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    "mymaster",
			SentinelAddrs: []string{"redis-sentinel-1:26379", "redis-sentinel-2:26380", "redis-sentinel-3:26381"},
		})
	} else {
		// for local debugging purposes
		rdb = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0, // use default DB
		})
	}

	return rdb
}
