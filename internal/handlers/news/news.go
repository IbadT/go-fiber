package news

import (
	"github.com/IbadT/go-fiber.git/database"
	"github.com/IbadT/go-fiber.git/internal/models"
	"github.com/gofiber/fiber/v2"
)

// CreateNews создает новую новость
func CreateNews(c *fiber.Ctx) error {
	news := new(models.News)
	if err := c.BodyParser(news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	// Получаем email пользователя из контекста (установлен middleware)
	userEmail := c.Locals("userEmail").(string)
	news.AuthorEmail = userEmail

	if err := database.DB.Create(&news).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create news",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   news,
	})
}

// GetNews получает список всех новостей
func GetNews(c *fiber.Ctx) error {
	var news []models.News
	if err := database.DB.Find(&news).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not fetch news",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   news,
	})
}

// GetNewsByID получает новость по ID
func GetNewsByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var news models.News

	if err := database.DB.First(&news, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "News not found",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   news,
	})
}

// UpdateNews обновляет новость
func UpdateNews(c *fiber.Ctx) error {
	id := c.Params("id")
	var news models.News

	// Проверяем существование новости
	if err := database.DB.First(&news, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "News not found",
		})
	}

	// Проверяем, является ли пользователь автором новости
	userEmail := c.Locals("userEmail").(string)
	if news.AuthorEmail != userEmail {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "You are not authorized to update this news",
		})
	}

	// Обновляем новость
	if err := c.BodyParser(&news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	if err := database.DB.Save(&news).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not update news",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   news,
	})
}

// DeleteNews удаляет новость
func DeleteNews(c *fiber.Ctx) error {
	id := c.Params("id")
	var news models.News

	// Проверяем существование новости
	if err := database.DB.First(&news, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "News not found",
		})
	}

	// Проверяем, является ли пользователь автором новости
	userEmail := c.Locals("userEmail").(string)
	if news.AuthorEmail != userEmail {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "You are not authorized to delete this news",
		})
	}

	if err := database.DB.Delete(&news).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not delete news",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "News deleted successfully",
	})
}
