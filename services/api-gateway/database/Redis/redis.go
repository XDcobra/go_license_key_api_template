package Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func ConnectionRedisDB() *redis.Client {
	var rdb *redis.Client
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// connect to redis
	rdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"gofiber-starter-stack-redis-sentinel-headless:26379"},
	})

	// check redis connection
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Redis connection error: %v", err)
	} else {
		log.Println("Connected to Redis")
	}

	return rdb
}
