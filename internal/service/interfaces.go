package service

import (
	"github.com/IbadT/go-fiber.git/internal/models"
	"github.com/google/uuid"
)

// UserService определяет интерфейс для бизнес-логики работы с пользователями
type UserService interface {
	// Register регистрирует нового пользователя
	Register(user *models.User) error
	// Login выполняет вход пользователя и возвращает JWT токен
	Login(email, password string) (string, error)
	// GetByID получает пользователя по ID
	GetByID(id uuid.UUID) (*models.User, error)
	// Update обновляет данные пользователя
	Update(user *models.User) error
	// Delete удаляет пользователя
	Delete(id uuid.UUID) error
}

// NewsService определяет интерфейс для бизнес-логики работы с новостями
type NewsService interface {
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

// CommentService определяет интерфейс для бизнес-логики работы с комментариями
type CommentService interface {
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
