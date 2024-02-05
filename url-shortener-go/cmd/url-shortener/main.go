package main

import (
	"fmt"
	"log"
	"os"
	"url-shortener-go/config"
	"url-shortener-go/modules/application"
	"url-shortener-go/modules/event"
	"url-shortener-go/modules/shorturl"
	"url-shortener-go/storage"
	"url-shortener-go/utils"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(prefix string, app *fiber.App) {
	logger := utils.CreateLogger("Main")
	application.SetupRoutes(prefix, app)
	prefix = prefix + "/api"
	shorturl.SetupRoutes(prefix, app)
	event.SetupRoutes(prefix, app)
	routes := app.GetRoutes()
	for _, route := range routes {
		logger.Info(fmt.Sprintf("Route Method: %s, Path: %s", route.Method, route.Path))
	}
}

func main() {
	config.LoadConfig()

	logger := utils.CreateLogger("Main")

	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	user := os.Getenv("DB_USER")
	fmt.Printf("DB_USER: %s\n", user)

	err := storage.ConnectDB()
	if err != nil {
		logger.Error(err, "Error connecting to database")
		panic("Error connecting to database")
	}

	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, // 50 MB in bytes
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return c.Next()
	})

	allowedOrigins := utils.Getenv("ALLOWED_ORIGINS", "*")
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	API_PREFIX := "/app"
	utils.InitTokenCounter()
	SetupRoutes(API_PREFIX, app)

	port := config.GetConfigValue("PORT")
	if port == "" {
		port = "8000" // Default port if not provided in .env
		log.Printf("PORT environment variable not set. Defaulting to %s", port)
	}

	log.Fatal(app.Listen(":" + port))
}
