package shorturl

import (
	"url-shortener-go/storage"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(prefix string, app *fiber.App) {
	shortUrlService := &ShortUrlService{DB: storage.DB}

	shortUrlController := InitShortUrlController(shortUrlService)

	urlgen := app.Group(prefix + "/urls")
	urlgen.Post("/create", shortUrlController.CreateShortUrl)
	urlgen.Get("/:shortUrl", shortUrlController.GetOriginalUrl)
	urlgen.Get("/", shortUrlController.GetShortUrls)

	app.Get("/:shortUrl", shortUrlController.RedirectToOriginalUrl)

}
