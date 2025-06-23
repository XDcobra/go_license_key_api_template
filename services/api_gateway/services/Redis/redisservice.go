package Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func GetKeyValues(rdb *redis.Client, key string) ([]string, error) {
	var keyValue []string
	var redisErr error

	ctx := context.Background()
	keyValue, redisErr = rdb.LRange(ctx, key, 0, -1).Result()

	return keyValue, redisErr
}

func SetKeyValue(rdb *redis.Client, key string, value string) error {
	ctx := context.Background()
	redisErr := rdb.LPush(ctx, key, value).Err()

	return redisErr
}
