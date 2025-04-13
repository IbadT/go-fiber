package models

// Импортируем необходимые пакеты
import (
	// Импорт GORM для работы с базой данных
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User представляет модель пользователя в системе
type User struct {
	// Встраиваем базовую модель GORM, которая добавляет поля:
	// ID        uint `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	gorm.Model

	ID        uuid.UUID `gorm:"type:uuid" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"` // `-` означает, что поле не будет сериализовано в JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Имя пользователя
	Name string `json:"name"`

	// // Дата последнего входа пользователя
	// LastLogin time.Time `json:"lastLogin"`
}

// UserLoginRequest представляет запрос на вход пользователя
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserRegisterRequest представляет запрос на регистрацию пользователя
type UserRegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// проверка пароля при входе
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
