package router

import (
	prom "github.com/XDcobra/gofiber-starter-stack/prometheus"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"os"
)

func CreateServer() *fiber.App {
	app := fiber.New()

	prometheus := prom.New("api_gateway")
	prometheus.RegisterAt(app, "/metrics", basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("METRICS_USER"): os.Getenv("METRICS_PASS"),
		},
	}))

	app.Use(prometheus.Middleware)

	return app
}
