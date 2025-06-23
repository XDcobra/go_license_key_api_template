package router

import (
	dummycontroller "github.com/XDcobra/gofiber-starter-stack/controller/DummyController"
	mysqlcontroller "github.com/XDcobra/gofiber-starter-stack/controller/MySQLController"
	rediscontroller "github.com/XDcobra/gofiber-starter-stack/controller/RedisController"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router *fiber.App, redisController *rediscontroller.RedisController, dummyController *dummycontroller.DummyController, mysqlController *mysqlcontroller.MySQLController) *fiber.App {
	// DummyController Endpoints
	router.Get("/", dummyController.DummyControllerPing)

	// RedisController Endpoints
	router.Get("/redis/ping", redisController.RedisControllerPing)
	router.Get("/redis/get", redisController.RedisControllerGet)
	router.Post("/redis/post", redisController.RedisControllerPost)

	// MySQLController Endpoints
	router.Get("/mysql/get/:id", mysqlController.MySQLControllerGet)
	router.Post("/mysql/post", mysqlController.MySQLControllerPost)

	return router
}
