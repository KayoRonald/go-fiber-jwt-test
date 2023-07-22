package main

import (
	"time"

	"github.com/KayoRonald/go-fiber-jwt-test/database"
	"github.com/KayoRonald/go-fiber-jwt-test/middleware"
	"github.com/KayoRonald/go-fiber-jwt-test/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/golang-jwt/jwt"
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
		user := []models.User{}
		result := database.Database.Db.Find(&user)
		if result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Nenhum usuário cadastrado",
				"status":  "err",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": user,
		})
	})
	//Private ROUTER
	app.Get("/me", middleware.VerifyToken,func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to router private",
		})
	})
	app.Get("/logout", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		c.ClearCookie()
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Logout",
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
			Name:     user.Name,
			Email:    user.Email,
			Password: string(hash),
		}
		result := database.Database.Db.Create(&createUser)
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Email já cadastrado! Tente outro",
				"status":  "err",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": &createUser,
			"status":  "sucess",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		playload := new(models.User)
		if err := c.BodyParser(playload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
				"status":  "err",
			})
		}
		var user models.User
		result := database.Database.Db.Where("email = ?", playload.Email).First(&user)
		if result.Error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Not exist User",
				"status":  "err",
			})
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(playload.Password))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid password! No match :/",
				"status":  "err",
			})
		}
		claims := jwt.MapClaims{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 2).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("1122222"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
				"status":  "sucess",
			})
		}
		c.Set("Authorization", tokenString)
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  time.Now().Add(2 * time.Hour),
			HTTPOnly: true,
			SameSite: "lax",
			Domain:   c.Hostname(),
		})
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": tokenString,
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
