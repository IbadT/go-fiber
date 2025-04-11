package router

import (
	userRoutes "github.com/IbadT/go-fiber.git/internal/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	userRoutes.SetupUserRoutes(api)
}
