package main

import (

	"github.com/KayoRonald/go-fiber-jwt-test/database"
	"github.com/KayoRonald/go-fiber-jwt-test/middleware"
	"github.com/KayoRonald/go-fiber-jwt-test/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	// "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(middleware.CorsMiddleware())
	app.Use(middleware.Limiter())
  // Conect Database
  database.ConnectDB()
  // metrics
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics My API"}))
	// Public ROUTER
	app.Get("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM. Router public",
		})
	})
	//Private ROUTER
	app.Get("/me", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to router private",
		})
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
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
				"status":  "err",
			})
		}
    createUser := models.User{
      Name: user.Name,
      Email: user.Email,
      Password: string(hash),
    }
    database.Database.Db.Create(&createUser)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": &createUser,
			"status":  "sucess",
		})
	})

	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
			"status":  "err",
		})
	})
	app.Listen(":3000")
}
