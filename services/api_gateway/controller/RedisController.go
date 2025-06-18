package controller

import (
	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
)

type RedisController struct {
	rdb *redis.Client
}

func NewRedisController(rdb *redis.Client) *RedisController {
	return &RedisController{
		rdb: rdb,
	}
}

func (n *RedisController) RedisControllerPing(c *fiber.Ctx) error {
	return c.SendString("Redis Controller Pong")
}
