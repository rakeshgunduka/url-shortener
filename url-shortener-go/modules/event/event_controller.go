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

// InitShortUrlController initializes the EventsController with the given EventService.
func InitShortUrlController(urlGenService *EventService) *EventsController {
	logger := utils.CreateLogger("InitUrlGenController")
	return &EventsController{urlGenService, logger}
}

// StoreEvents stores the events in the analytics service.
// @Summary Store events
// @Description Store the events in the analytics service
// @ID store-events
// @Accept  json
// @Produce json
// @Param   name path string true "Event name"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /app/api/events/{name} [post]
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

// GetEvents returns the events from the analytics service.
// @Summary Get events
// @Description Get the events from the analytics service
// @ID get-events
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /app/api/events/ [get]
func (ac *EventsController) GetEvents(c *fiber.Ctx) error {
	events := ac.analyticsService.GetAnalytics()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"events": events,
	})
}
