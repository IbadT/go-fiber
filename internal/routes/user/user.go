package userRoutes

import (
	userHandler "github.com/IbadT/go-fiber.git/internal/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {
	router.Post("/register", userHandler.RegisterUser)

	router.Post("/login", userHandler.LoginUser)

	router.Get("/:userId", userHandler.GetUser)
}
