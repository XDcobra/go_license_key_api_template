package DummyController

import (
	"github.com/gofiber/fiber/v2"
)

type DummyController struct {
}

func NewDummyController() *DummyController {
	return &DummyController{}
}

func (n *DummyController) DummyControllerPing(c *fiber.Ctx) error {
	return c.SendString("Pong")
}
