package application

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(prefix string, app *fiber.App) {
	appController := InitAppController()

	app.Get(prefix+"/hc", appController.HealthCheck)
}
