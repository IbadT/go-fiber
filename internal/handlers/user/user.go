package userHandler

import (
	"time"

	"github.com/IbadT/go-fiber.git/config"
	"github.com/IbadT/go-fiber.git/database"
	"github.com/IbadT/go-fiber.git/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GetUser godoc
// @Summary Get user by ID
// @Description Get user information by ID
// @Tags users
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Security ApiKeyAuth
// @Router /{userId} [get]
func GetUser(c *fiber.Ctx) error {
	// xh http://localhost:8000/api/f7acf2fd-0b3b-444c-be07-d6008fc4b983
	db := database.DB
	id := c.Params("userId")
	var user models.User

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(map[string]interface{}{"status": "error", "message": "User is not founded", "data": nil})
	}

	return c.JSON(map[string]interface{}{"status": "success", "message": "User founded", "data": user})
}

type RegisterRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"securepassword"`
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User registration data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register [post]
func RegisterUser(c *fiber.Ctx) error {
	// xh post http://localhost:8000/api/register \
	//     email="user@example.com" \
	//     password="securepassword"
	db := database.DB

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(map[string]interface{}{"status": "error", "message": "Review yout input", "data": err})
	}

	if err := user.HashPassword(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": "Ошибка хеширования пароля"})
	}

	user.ID = uuid.New()
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(map[string]interface{}{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.JSON(map[string]interface{}{"status": "success", "message": "Successfull register user", "data": user})
}

type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"securepassword"`
}

// LoginUser godoc
// @Summary Login user
// @Description Authenticate user and return token
// @Tags users
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "User login credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /login [post]
func LoginUser(c *fiber.Ctx) error {
	// xh post http://localhost:8000/api/login \
	//     email="user@example.com" \
	//     password="securepassword"
	db := database.DB

	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(map[string]interface{}{"status": "error", "message": "Review yout input", "data": err})
	}

	var dbUser models.User

	if err := db.Find(&dbUser, "email = ?", user.Email).Error; err != nil {
		return c.Status(500).JSON(map[string]interface{}{"status": "error", "message": "Invalid Email", "data": err})
	}

	if !dbUser.CheckPassword(user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{"error": "Неверный пароль"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": dbUser.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(config.Config("JWT_SECRET_WORD")))

	// Устанавливаем cookie
	c.Cookie(&fiber.Cookie{
		Name:    "access_token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return c.JSON(map[string]interface{}{"message": "Вход выполнен", "token": tokenString})
}
