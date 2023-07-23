package middleware

import (
	"fmt"
	"strings"

	"github.com/KayoRonald/go-fiber-jwt-test/database"
	"github.com/KayoRonald/go-fiber-jwt-test/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type ClaimsD struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func VerifyToken(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": "You are not logged in",
		})
	}
	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte("1122222"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("invalidate token: %v", err),
		})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "fail",
			"message": "invalid token claim",
		})
	}
  var user models.User
  database.Database.Db.Where("id = ?", fmt.Sprint(claims["id"])).First(&user)
  if user.ID != claims["id"] {
    return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "fail",
			"message": "invalid user",
		})
  }
	ClaimsD := ClaimsD{
		ID: claims["id"].(string),
		Name: claims["name"].(string),
		Email: claims["email"].(string),
	}
	fmt.Printf("User authenticated: %#v \n", ClaimsD)
	return c.Next()
}
