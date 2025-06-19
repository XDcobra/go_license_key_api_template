package router

import (
	prom "github.com/XDcobra/go_license_key_api_template/prometheus"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func CreateServer() *fiber.App {
	app := fiber.New()

	prometheus := prom.New("api_gateway")
	prometheus.RegisterAt(app, "/metrics", basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testuser": "testpw",
		},
	}))

	app.Use(prometheus.Middleware)

	return app
}
