package shorturl

import (
	"strings"
	"time"
	"url-shortener-go/storage"
	"url-shortener-go/utils"

	"github.com/gofiber/fiber/v2"
)

type ShortUrlController struct {
	shortUrlService *ShortUrlService
	logger          *utils.Logger
}

func InitShortUrlController(urlGenService *ShortUrlService) *ShortUrlController {
	logger := utils.CreateLogger("InitUrlGenController")
	return &ShortUrlController{urlGenService, logger}
}

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

func (sc *ShortUrlController) RedirectToOriginalUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")
	redisKey := "short-url-" + shortUrl

	cachedResponse, _ := storage.RedisGet(redisKey)
	if cachedResponse != nil {
		originalURL, ok := cachedResponse["OriginalUrl"].(string)
		if !ok {
			return c.Redirect("http://localhost:3000/404")
		}
		return c.Redirect(originalURL, fiber.StatusMovedPermanently)
	}

	urlAlias, err := sc.shortUrlService.GetOriginalUrl(shortUrl)

	if err != nil {
		return c.Redirect("http://localhost:3000/404")
	}

	if urlAlias.Expiry != (time.Time{}) && time.Now().After(urlAlias.Expiry) {
		return c.Redirect("http://localhost:3000/404")
	}

	storage.RedisSet(redisKey, map[string]interface{}{"OriginalUrl": urlAlias.URL, "Expiry": urlAlias.Expiry}, time.Hour*1)

	return c.Redirect(urlAlias.URL, fiber.StatusMovedPermanently)
}

func (uc *ShortUrlController) GetShortUrls(c *fiber.Ctx) error {
	urls, _ := uc.shortUrlService.GetShortUrls()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"urls":   urls,
	})
}
