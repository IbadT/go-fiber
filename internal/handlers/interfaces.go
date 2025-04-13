package handlers

import "github.com/gofiber/fiber/v2"

// UserHandler определяет интерфейс для обработки HTTP-запросов, связанных с пользователями
type UserHandler interface {
	// RegisterUser обрабатывает запрос на регистрацию нового пользователя
	RegisterUser(c *fiber.Ctx) error
	// LoginUser обрабатывает запрос на вход пользователя
	LoginUser(c *fiber.Ctx) error
	// GetUser обрабатывает запрос на получение информации о пользователе
	GetUser(c *fiber.Ctx) error
}
