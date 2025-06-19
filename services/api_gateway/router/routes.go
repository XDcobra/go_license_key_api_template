package router

import (
	dummycontroller "github.com/XDcobra/go_license_key_api_template/controller/DummyController"
	rediscontroller "github.com/XDcobra/go_license_key_api_template/controller/RedisController"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router *fiber.App, redisController *rediscontroller.RedisController, dummyController *dummycontroller.DummyController) *fiber.App {
	// DummyController Endpoints
	router.Get("/", dummyController.DummyControllerPing)

	// RedisController Endpoints
	router.Get("/redis/ping", redisController.RedisControllerPing)
	router.Get("/redis/get", redisController.RedisControllerGet)
	router.Post("/redis/post", redisController.RedisControllerPost)

	return router
}
