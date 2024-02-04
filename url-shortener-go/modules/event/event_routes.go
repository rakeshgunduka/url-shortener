package event

import (
	"url-shortener-go/storage"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(prefix string, app *fiber.App) {
	eventService := &EventService{DB: storage.DB}

	eventController := InitShortUrlController(eventService)

	event := app.Group(prefix + "/events")
	event.Get("/", eventController.GetEvents)
	event.Post("/:name", eventController.StoreEvents)
}
