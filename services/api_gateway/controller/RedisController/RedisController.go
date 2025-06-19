package RedisController

import (
	"github.com/XDcobra/go_license_key_api_template/model"
	"github.com/XDcobra/go_license_key_api_template/services"
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
	var redisPayloadModel model.RedisPayloadModel

	// get request parameter
	err := c.BodyParser(&redisPayloadModel)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.RedisGetResponseModel{Errors: err.Error()})
	}

	// get key value from redis
	keyValue, redisErr := services.GetKeyValues(n.rdb, redisPayloadModel.Key)
	if redisErr != nil {
		return c.Status(http.StatusBadRequest).JSON(model.RedisGetResponseModel{Errors: redisErr.Error()})
	}

	// return key value
	return c.Status(http.StatusOK).JSON(model.RedisGetResponseModel{Key: redisPayloadModel.Key, Values: keyValue})
}

func (n *RedisController) RedisControllerPost(c *fiber.Ctx) error {
	var redisPayloadModel model.RedisPayloadModel

	// get request parameter
	err := c.BodyParser(&redisPayloadModel)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.RedisPostResponseModel{Errors: "Body parser: " + err.Error()})
	}

	// set key value in redis
	redisErr := services.SetKeyValue(n.rdb, redisPayloadModel.Key, redisPayloadModel.Value)
	if redisErr != nil {
		return c.Status(http.StatusBadRequest).JSON(model.RedisPostResponseModel{Errors: "Setting in rdb: " + redisErr.Error()})
	}

	// return ok
	return c.Status(http.StatusOK).JSON(model.RedisPostResponseModel{Key: redisPayloadModel.Key, Value: redisPayloadModel.Value})
}
