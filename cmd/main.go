package main

import (
	"fmt"
	"log"

	"github.com/IbadT/go-fiber.git/config"
	"github.com/IbadT/go-fiber.git/database"
	_ "github.com/IbadT/go-fiber.git/internal/docs" // Подключение Swagger
	"github.com/IbadT/go-fiber.git/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	swagger "github.com/swaggo/fiber-swagger"
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
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file %s", err.Error())
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte("Hello"))
	})

	env := config.Config("APP_ENV")
	port := config.Config("PORT")

	if env == "development" {
		// Swagger configuration
		app.Get("/swagger/*", swagger.WrapHandler)
	}

	log.Fatal(app.Listen(port))
}

// swag init -g cmd/main.go --output internal/docs
