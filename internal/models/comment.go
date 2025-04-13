package models

// Импортируем необходимые пакеты
import (
	// Импорт GORM для работы с базой данных
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Comment представляет модель комментария в системе
type Comment struct {
	// Встраиваем базовую модель GORM, которая добавляет поля:
	// ID        uint `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`

	// Содержание комментария
	Content string `json:"content" gorm:"not null"`

	// ID новости, к которой относится комментарий (связь с моделью News)
	NewsID uint `json:"newsId" gorm:"not null"`

	// Email автора комментария (связь с моделью User)
	AuthorEmail string `json:"authorEmail" gorm:"not null"`
}
