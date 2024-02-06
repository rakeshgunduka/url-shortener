package shorturl

import (
	"strings"
	"time"
	"url-shortener-go/config"
	"url-shortener-go/modules/event"
	"url-shortener-go/storage"
	"url-shortener-go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/ua-parser/uap-go/uaparser"
)

type ShortUrlController struct {
	eventService    *event.EventService
	shortUrlService *ShortUrlService
	logger          *utils.Logger
}

func InitShortUrlController(eventService *event.EventService, urlGenService *ShortUrlService) *ShortUrlController {
	logger := utils.CreateLogger("InitUrlGenController")
	return &ShortUrlController{eventService, urlGenService, logger}
}

// CreateShortUrl creates a short URL for the given long URL.
// @Summary Create a short URL
// @Description Create a short URL for the given long URL
// @ID create-short-url
// @Accept  json
// @Produce json
// @Param   url body string true "Long URL"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /app/api/urls/create [post]
func (uc *ShortUrlController) CreateShortUrl(c *fiber.Ctx) error {

	var requestBody struct {
		LongURL string `json:"url"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid payload",
		})
	}

	longUrl := requestBody.LongURL

	if !strings.HasPrefix(longUrl, "http://") && !strings.HasPrefix(longUrl, "https://") {
		longUrl = "http://" + longUrl
	}

	shortUrl, err := uc.shortUrlService.GetShortUrl(longUrl)
	if shortUrl != "" && err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "URL already exists",
			"status":  "ok",
			"alias":   shortUrl,
		})
	}

	shortUrl = uc.shortUrlService.GenerateUrlAlias(longUrl)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"alias":  shortUrl,
	})
}

// GetOriginalUrl retrieves the original URL for a given short URL.
// @Summary Retrieve the original URL
// @Description Retrieve the original URL for a given short URL
// @ID get-original-url
// @Accept  json
// @Produce json
// @Param   shortUrl path string true "Short URL"
// @Success 200 {object} string
// @Failure 404 {object} string
// @Router /app/api/urls/{shortUrl} [get]
func (sc *ShortUrlController) GetOriginalUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	redisKey := "short-url-" + shortUrl

	cachedResponse, _ := storage.RedisGet(redisKey)
	if cachedResponse != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":      "ok",
			"cached":      true,
			"originalUrl": cachedResponse["OriginalUrl"],
			"expiry":      cachedResponse["Expiry"],
		})
	}

	urlAlias, err := sc.shortUrlService.GetOriginalUrl(shortUrl)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	if urlAlias.Expiry != (time.Time{}) && time.Now().After(urlAlias.Expiry) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL expired",
		})
	}

	storage.RedisSet(redisKey, map[string]interface{}{"OriginalUrl": urlAlias.URL, "Expiry": urlAlias.Expiry}, time.Hour*1)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "ok",
		"originalUrl": urlAlias.URL,
	})
}

// RedirectToOriginalUrl redirects to the original URL for a given short URL.
// @Summary Redirect to the original URL
// @Description Redirect to the original URL for a given short URL
// @ID redirect-original-url
// @Accept  json
// @Produce json
// @Param   shortUrl path string true "Short URL"
// @Success 301 {object} string
// @Failure 404 {object} string
// @Router /app/{shortUrl} [get]
func (sc *ShortUrlController) RedirectToOriginalUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	redisKey := "short-url-" + shortUrl

	parser := uaparser.NewFromSaved()
	eventName := event.EventName(c.Params("name"))
	userAgent := c.Get("User-Agent")
	client := parser.Parse(userAgent)

	go sc.eventService.StoreAnalytics(eventName, userAgent, client.Os.Family, client.Device.Brand, client.Device.Family, client.Device.Model)

	cachedResponse, _ := storage.RedisGet(redisKey)
	if cachedResponse != nil {
		originalURL, ok := cachedResponse["OriginalUrl"].(string)
		if !ok {
			return c.Redirect(config.GetConfigValue("REACT_APP_URL") + "/not-found")
		}
		return c.Redirect(originalURL, fiber.StatusMovedPermanently)
	}

	urlAlias, err := sc.shortUrlService.GetOriginalUrl(shortUrl)

	if err != nil {
		return c.Redirect(config.GetConfigValue("REACT_APP_URL") + "/not-found")
	}

	if urlAlias.Expiry != (time.Time{}) && time.Now().After(urlAlias.Expiry) {
		return c.Redirect(config.GetConfigValue("REACT_APP_URL") + "/not-found")
	}

	storage.RedisSet(redisKey, map[string]interface{}{"OriginalUrl": urlAlias.URL, "Expiry": urlAlias.Expiry}, time.Hour*1)

	return c.Redirect(urlAlias.URL, fiber.StatusMovedPermanently)
}

// GetShortUrls retrieves all the short URLs.
// @Summary Retrieve all short URLs
// @Description Retrieve all short URLs
// @ID get-short-urls
// @Accept  json
// @Produce json
// @Success 200 {object} string
// @Router /app/api/urls/ [get]
func (uc *ShortUrlController) GetShortUrls(c *fiber.Ctx) error {
	urls, _ := uc.shortUrlService.GetShortUrls()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"urls":   urls,
	})
}
