package middleware

// Импортируем необходимые пакеты
import (
	// Пакет для работы со строками
	"strings"

	// Импорт конфигурации приложения
	"github.com/IbadT/go-fiber.git/config"
	// Импорт фреймворка Fiber
	"github.com/gofiber/fiber/v2"
	// Импорт пакета для работы с JWT токенами
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware создает middleware для проверки JWT токена
// Возвращает функцию-обработчик Fiber
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Получаем значение заголовка Authorization
		authHeader := c.Get("Authorization")
		var tokenString string

		// Проверяем наличие токена в заголовке
		if authHeader != "" {
			// Разбиваем заголовок на части (Bearer и сам токен)
			parts := strings.Split(authHeader, " ")
			// Проверяем формат заголовка
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":  "error",
					"message": "Invalid authorization header format",
				})
			}
			// Извлекаем сам токен
			tokenString = parts[1]
		} else {
			// Если токена нет в заголовке, пробуем получить из cookie
			tokenString = c.Cookies("jwt")
		}

		// Проверяем, что токен был найден
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Authorization token is required",
			})
		}

		// Парсим и проверяем токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Проверяем, что используется правильный алгоритм подписи
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			// Возвращаем секретный ключ для проверки подписи
			return []byte(config.Config("JWT_SECRET_WORD")), nil
		})

		// Проверяем ошибки парсинга токена
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid or expired token",
			})
		}

		// Проверяем валидность токена
		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid token",
			})
		}

		// Извлекаем данные из токена
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid token claims",
			})
		}

		// Сохраняем email пользователя в контексте для использования в обработчиках
		c.Locals("userEmail", claims["email"])

		// Продолжаем выполнение запроса
		return c.Next()
	}
}
