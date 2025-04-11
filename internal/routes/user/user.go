package userRoutes

import (
	userHandler "github.com/IbadT/go-fiber.git/internal/handlers/user"
	"github.com/gofiber/fiber/v3"
)

func SetupNoteRoutes(router fiber.Router) {
	// user := router.Group("/api")

	router.Post("/register", userHandler.RegisterUser)

	router.Post("/login", userHandler.LoginUser)

	router.Get("/:userId", userHandler.GetUser)
}
