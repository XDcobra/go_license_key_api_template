package RedisController

import (
	RedisModels "github.com/XDcobra/gofiber-starter-stack/model/Redis"
	RedisService "github.com/XDcobra/gofiber-starter-stack/services/Redis"
	"github.com/redis/go-redis/v9"
	"net/http"

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

func (n *RedisController) RedisControllerGet(c *fiber.Ctx) error {
	var redisPayloadModel RedisModels.RedisPayloadModel

	// get request parameter
	err := c.BodyParser(&redisPayloadModel)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(RedisModels.RedisGetResponseModel{Errors: err.Error()})
	}

	// get key value from redis
	keyValue, redisErr := RedisService.GetKeyValues(n.rdb, redisPayloadModel.Key)
	if redisErr != nil {
		return c.Status(http.StatusBadRequest).JSON(RedisModels.RedisGetResponseModel{Errors: redisErr.Error()})
	}

	// return key value
	return c.Status(http.StatusOK).JSON(RedisModels.RedisGetResponseModel{Key: redisPayloadModel.Key, Values: keyValue})
}

func (n *RedisController) RedisControllerPost(c *fiber.Ctx) error {
	var redisPayloadModel RedisModels.RedisPayloadModel

	// get request parameter
	err := c.BodyParser(&redisPayloadModel)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(RedisModels.RedisPostResponseModel{Errors: "Body parser: " + err.Error()})
	}

	// set key value in redis
	redisErr := RedisService.SetKeyValue(n.rdb, redisPayloadModel.Key, redisPayloadModel.Value)
	if redisErr != nil {
		return c.Status(http.StatusBadRequest).JSON(RedisModels.RedisPostResponseModel{Errors: "Setting in rdb: " + redisErr.Error()})
	}

	// return ok
	return c.Status(http.StatusOK).JSON(RedisModels.RedisPostResponseModel{Key: redisPayloadModel.Key, Value: redisPayloadModel.Value})
}
