package auth

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		tokenString := c.Get("Authorization")

		tokenString = strings.TrimSpace(tokenString)
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimSpace(tokenString[7:])
		}

		if tokenString == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "token requerido",
			})
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "secret"
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "token invalido",
			})
		}

		return c.Next()
	}
}
