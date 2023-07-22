package main

import (
	"time"

	"github.com/KayoRonald/go-fiber-jwt-test/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)
func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://*, https://*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
    AllowCredentials: false,
	}))

	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics My API"}))
	
  app.Get("/", func(c *fiber.Ctx) error {
    c.Accepts("application/json")
		return c.SendString("Hello, World!")
	})
	app.Post("/sinup", func(c *fiber.Ctx) error {
    c.Accepts("application/json")
    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
        "message": err.Error(),
        "status":  "err",
      })
    }
    return c.SendString("")
	})

	app.Listen(":3000")
}
