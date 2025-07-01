package router

import (
	prom "github.com/XDcobra/gofiber-starter-stack/prometheus"
	"github.com/gofiber/fiber/v2"
)

func CreateServer() *fiber.App {
	app := fiber.New()

	prometheus := prom.New("api-gateway")

	prometheus.RegisterAt(app, "/metrics")
	/* in case you want to add basic auth
	prometheus.RegisterAt(app, "/metrics", basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("METRICS_USER"): os.Getenv("METRICS_PASS"),
		},
	}))
	*/

	app.Use(prometheus.Middleware)

	return app
}
