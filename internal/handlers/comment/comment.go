package comment

import (
	"github.com/IbadT/go-fiber.git/database"
	"github.com/IbadT/go-fiber.git/internal/models"
	"github.com/gofiber/fiber/v2"
)

// CreateComment создает новый комментарий
func CreateComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	// Получаем email пользователя из контекста (установлен middleware)
	userEmail := c.Locals("userEmail").(string)
	comment.AuthorEmail = userEmail

	if err := database.DB.Create(&comment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create comment",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   comment,
	})
}

// GetCommentsByNewsID получает все комментарии к новости
func GetCommentsByNewsID(c *fiber.Ctx) error {
	newsID := c.Params("newsId")
	var comments []models.Comment

	if err := database.DB.Where("news_id = ?", newsID).Find(&comments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not fetch comments",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   comments,
	})
}

// UpdateComment обновляет комментарий
func UpdateComment(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment

	// Проверяем существование комментария
	if err := database.DB.First(&comment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Comment not found",
		})
	}

	// Проверяем, является ли пользователь автором комментария
	userEmail := c.Locals("userEmail").(string)
	if comment.AuthorEmail != userEmail {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "You are not authorized to update this comment",
		})
	}

	// Обновляем комментарий
	if err := c.BodyParser(&comment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	if err := database.DB.Save(&comment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not update comment",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   comment,
	})
}

// DeleteComment удаляет комментарий
func DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	var comment models.Comment

	// Проверяем существование комментария
	if err := database.DB.First(&comment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Comment not found",
		})
	}

	// Проверяем, является ли пользователь автором комментария
	userEmail := c.Locals("userEmail").(string)
	if comment.AuthorEmail != userEmail {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "You are not authorized to delete this comment",
		})
	}

	if err := database.DB.Delete(&comment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not delete comment",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Comment deleted successfully",
	})
}
