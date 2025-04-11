package middleware

import (
	"strings"

	"github.com/IbadT/go-fiber.git/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware проверяет JWT токен в заголовке Authorization или в cookie
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Получаем токен из заголовка Authorization
		authHeader := c.Get("Authorization")
		var tokenString string

		if authHeader != "" {
			// Формат: "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  "error",
					"message": "Invalid authorization header format",
				})
			}
			tokenString = parts[1]
		} else {
			// Если токена нет в заголовке, пробуем получить из cookie
			tokenString = c.Cookies("jwt")
		}

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Authorization token is required",
			})
		}

		// Проверяем токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Проверяем, что используется правильный алгоритм
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(config.Config("JWT_SECRET_WORD")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid or expired token",
			})
		}

		// Проверяем, что токен валидный
		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid token",
			})
		}

		// Получаем данные из токена
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid token claims",
			})
		}

		// Сохраняем email пользователя в контексте для использования в обработчиках
		c.Locals("userEmail", claims["email"])

		return c.Next()
	}
}
