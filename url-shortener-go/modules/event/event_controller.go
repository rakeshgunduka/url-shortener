package event

import (
	"url-shortener-go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/ua-parser/uap-go/uaparser"
)

type EventsController struct {
	analyticsService *EventService
	logger           *utils.Logger
}

type EventName string

const (
	EventNameClick  EventName = "click"
	EventNameView   EventName = "view"
	EventNameSubmit EventName = "submit"
)

func InitShortUrlController(urlGenService *EventService) *EventsController {
	logger := utils.CreateLogger("InitUrlGenController")
	return &EventsController{urlGenService, logger}
}

func (ac *EventsController) StoreEvents(c *fiber.Ctx) error {
	parser := uaparser.NewFromSaved()
	eventName := EventName(c.Params("name"))
	userAgent := c.Get("User-Agent")
	client := parser.Parse(userAgent)

	go ac.analyticsService.StoreAnalytics(eventName, userAgent, client.Os.Family, client.Device.Brand, client.Device.Family, client.Device.Model)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func (ac *EventsController) GetEvents(c *fiber.Ctx) error {
	events := ac.analyticsService.GetAnalytics()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"events": events,
	})
}
