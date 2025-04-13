package models

// Импортируем необходимые пакеты
import (
	// Пакет для работы с временем
	"time"
	// Импорт GORM для работы с базой данных
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// News представляет модель новости в системе
type News struct {
	// Встраиваем базовую модель GORM, которая добавляет поля:
	// ID        uint `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`

	// Заголовок новости
	Title string `json:"title" gorm:"not null"`

	// Содержание новости
	Content string `json:"content" gorm:"not null"`

	// Email автора новости (связь с моделью User)
	AuthorEmail string `json:"authorEmail" gorm:"not null"`

	// Дата публикации новости
	PublishedAt time.Time `json:"publishedAt"`
}
