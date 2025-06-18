package router

import (
	rediscontroller "github.com/XDcobra/go_license_key_api_template/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router *fiber.App, rediscontroller *rediscontroller.RedisController) *fiber.App {
	// AdminController Endpoints
	router.Get("/redis/ping", rediscontroller.RedisControllerPing)

	return router
}
