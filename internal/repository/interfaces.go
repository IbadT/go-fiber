package repository

import (
	"github.com/IbadT/go-fiber.git/internal/models"
	"github.com/google/uuid"
)

// UserRepository определяет интерфейс для работы с пользователями в базе данных
type UserRepository interface {
	// Create создает нового пользователя
	Create(user *models.User) error
	// GetByID получает пользователя по ID
	GetByID(id uuid.UUID) (*models.User, error)
	// GetByEmail получает пользователя по email
	GetByEmail(email string) (*models.User, error)
	// Update обновляет данные пользователя
	Update(user *models.User) error
	// Delete удаляет пользователя
	Delete(id uuid.UUID) error
}

// NewsRepository определяет интерфейс для работы с новостями в базе данных
type NewsRepository interface {
	// Create создает новую новость
	Create(news *models.News) error
	// GetByID получает новость по ID
	GetByID(id uuid.UUID) (*models.News, error)
	// GetAll получает все новости
	GetAll() ([]models.News, error)
	// Update обновляет новость
	Update(news *models.News) error
	// Delete удаляет новость
	Delete(id uuid.UUID) error
	// GetByAuthor получает все новости автора
	GetByAuthor(authorEmail string) ([]models.News, error)
}

// CommentRepository определяет интерфейс для работы с комментариями в базе данных
type CommentRepository interface {
	// Create создает новый комментарий
	Create(comment *models.Comment) error
	// GetByID получает комментарий по ID
	GetByID(id uuid.UUID) (*models.Comment, error)
	// GetByNewsID получает все комментарии к новости
	GetByNewsID(newsID uuid.UUID) ([]models.Comment, error)
	// Update обновляет комментарий
	Update(comment *models.Comment) error
	// Delete удаляет комментарий
	Delete(id uuid.UUID) error
	// GetByAuthor получает все комментарии автора
	GetByAuthor(authorEmail string) ([]models.Comment, error)
}
