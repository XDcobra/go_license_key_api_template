package MySQLController

import (
	"fmt"
	MySQlModels "github.com/XDcobra/gofiber-starter-stack/model/MySQL"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type MySQLController struct {
	db *gorm.DB
}

func NewMySQLController(db *gorm.DB) *MySQLController {
	return &MySQLController{
		db: db,
	}
}

func (n *MySQLController) MySQLControllerPost(c *fiber.Ctx) error {
	var user MySQlModels.User

	// Parse request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Save user to DB
	if err := n.db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (n *MySQLController) MySQLControllerGet(c *fiber.Ctx) error {
	id := c.Params("id")
	var user MySQlModels.User

	if err := n.db.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("User with ID %s not found", id),
		})
	}

	return c.JSON(user)
}
