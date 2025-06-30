package router

import (
	dummycontroller "github.com/XDcobra/gofiber-starter-stack/controller/DummyController"
	mysqlcontroller "github.com/XDcobra/gofiber-starter-stack/controller/MySQLController"
	rediscontroller "github.com/XDcobra/gofiber-starter-stack/controller/RedisController"
	_ "github.com/XDcobra/gofiber-starter-stack/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
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

	// Swagger
	router.Get("/api/swagger/*", swagger.New(swagger.Config{
		URL: "/api/swagger/doc.json",
	}))

	return router
}
