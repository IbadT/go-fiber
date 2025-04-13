package repository

import (
	"github.com/IbadT/go-fiber.git/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepositoryImpl реализует интерфейс UserRepository
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository создает новый экземпляр UserRepositoryImpl
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// Create создает нового пользователя
func (r *UserRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetByID получает пользователя по ID
func (r *UserRepositoryImpl) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail получает пользователя по email
func (r *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update обновляет данные пользователя
func (r *UserRepositoryImpl) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete удаляет пользователя
func (r *UserRepositoryImpl) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}
