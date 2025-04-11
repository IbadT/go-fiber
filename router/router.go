package router

import (
	userHandler "github.com/IbadT/go-fiber.git/internal/handlers/user"
	"github.com/IbadT/go-fiber.git/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Публичные маршруты (без авторизации)
	api.Post("/register", userHandler.RegisterUser)
	api.Post("/login", userHandler.LoginUser)

	// Защищенные маршруты (требуют авторизации)
	protected := api.Group("/", middleware.AuthMiddleware())
	protected.Get("/:userId", userHandler.GetUser)
}
