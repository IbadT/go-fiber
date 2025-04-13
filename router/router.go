package router

// Импортируем необходимые пакеты
import (
	// Импорт обработчика комментариев с алиасом commentHandler
	commentHandler "github.com/IbadT/go-fiber.git/internal/handlers/comment"
	// Импорт обработчика новостей с алиасом newsHandler
	newsHandler "github.com/IbadT/go-fiber.git/internal/handlers/news"
	// Импорт обработчика пользователей с алиасом userHandler
	userHandler "github.com/IbadT/go-fiber.git/internal/handlers/user"
	// Импорт middleware для авторизации
	"github.com/IbadT/go-fiber.git/internal/middleware"

	// Импорт основного фреймворка Fiber
	"github.com/gofiber/fiber/v2"
	// Импорт middleware для логирования
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes настраивает все маршруты приложения
// Принимает экземпляр приложения Fiber в качестве параметра
func SetupRoutes(app *fiber.App) {
	// Создаем группу маршрутов с префиксом /api и добавляем middleware логирования
	api := app.Group("/api", logger.New())

	// Публичные маршруты (без авторизации)
	// Регистрация нового пользователя
	api.Post("/register", userHandler.RegisterUser)
	// Вход пользователя в систему
	api.Post("/login", userHandler.LoginUser)
	// Получение списка всех новостей
	api.Get("/news", newsHandler.GetNews)
	// Получение конкретной новости по ID
	api.Get("/news/:id", newsHandler.GetNewsByID)

	// Создаем группу защищенных маршрутов с middleware авторизации
	protected := api.Group("/", middleware.AuthMiddleware())

	// Маршруты для пользователей
	// Получение информации о пользователе по ID
	protected.Get("/:userId", userHandler.GetUser)

	// Маршруты для новостей
	// Создание новой новости
	protected.Post("/news", newsHandler.CreateNews)
	// Обновление существующей новости по ID
	protected.Put("/news/:id", newsHandler.UpdateNews)
	// Удаление новости по ID
	protected.Delete("/news/:id", newsHandler.DeleteNews)

	// Маршруты для комментариев
	// Создание нового комментария к новости
	protected.Post("/news/:newsId/comments", commentHandler.CreateComment)
	// Получение всех комментариев к новости
	protected.Get("/news/:newsId/comments", commentHandler.GetCommentsByNewsID)
	// Обновление существующего комментария по ID
	protected.Put("/comments/:id", commentHandler.UpdateComment)
	// Удаление комментария по ID
	protected.Delete("/comments/:id", commentHandler.DeleteComment)
}
