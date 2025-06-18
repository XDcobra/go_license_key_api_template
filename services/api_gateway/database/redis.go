package database

import (
	"github.com/redis/go-redis/v9"
)

func ConnectionRedisDB() *redis.Client {
	var rdb *redis.Client

	rdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"redis-sentinel-1:26379", "redis-sentinel-2:26380", "redis-sentinel-3:26381"},
	})

	return rdb
}
