package application

import (
	"github.com/gofiber/fiber/v2"
)

type AppController struct {
}

func InitAppController() *AppController {
	return &AppController{}
}

func (ac *AppController) HealthCheck(c *fiber.Ctx) error {
	return c.JSON("PONG")
}
