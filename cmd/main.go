package main

import (
	"fmt"
	"log"

	"github.com/IbadT/go-fiber.git/database"
	_ "github.com/IbadT/go-fiber.git/internal/docs" // Подключение Swagger
	"github.com/IbadT/go-fiber.git/router"
	httpSwagger "github.com/swaggo/http-swagger"

	// swagger "github.com/arsmn/fiber-swagger/v2" // Альтернативный пакет
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v3"

	"github.com/joho/godotenv"
)

// @title Fiber User API
// @version 1.0
// @description This is a sample server for a user management API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /api
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file %s", err.Error())
	}

	app := fiber.New()

	// app.Use(swagger.New())

	// app.Get("/swagger/*", fiberSwagger.WrapHandler)
	// Конвертируем стандартный http.Handler в обработчик Fiber
	swaggerHandler := adaptor.HTTPMiddleware(
		httpSwagger.Handler(
			httpSwagger.URL("/swagger/doc.json"),
		),
	)

	// Настройка маршрута для Swagger UI
	app.Get("/swagger/*", swaggerHandler)

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Get("/", func(c fiber.Ctx) error {
		return c.Send([]byte("Hello"))
	})

	log.Fatal(app.Listen(":8000"))
}

// swag init -g cmd/main.go --output internal/docs
