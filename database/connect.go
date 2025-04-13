package database

// Импортируем необходимые пакеты
import (
	// Пакет для форматирования строк
	"fmt"
	// Пакет для логирования
	"log"
	// Пакет для конвертации строк в числа
	"strconv"

	// Импорт конфигурации приложения
	"github.com/IbadT/go-fiber.git/config"
	// Импорт моделей данных
	"github.com/IbadT/go-fiber.git/internal/models"
	// Импорт драйвера PostgreSQL для GORM
	"gorm.io/driver/postgres"
	// Импорт ORM GORM
	"gorm.io/gorm"
)

// DB - глобальная переменная для хранения подключения к базе данных
var DB *gorm.DB

// ConnectDB устанавливает соединение с базой данных и выполняет миграции
func ConnectDB() {
	var err error
	// Получаем порт из конфигурации
	p := config.Config("DB_PORT")
	// Конвертируем строку в число
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		// В случае ошибки используем порт по умолчанию
		log.Printf("Error parsing port: %v, using default 5432\n", err)
		port = 5432
	}

	// Получаем параметры подключения из конфигурации
	host := config.Config("DB_HOST")
	user := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	dbname := config.Config("DB_NAME")

	// Логируем параметры подключения (без пароля)
	log.Printf("Connecting to PostgreSQL on %s:%d as %s\n", host, port, user)

	// Формируем строку подключения к базе данных
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Устанавливаем соединение с базой данных
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Printf("Failed to connect to database: %v\n", err)
		panic("Failed to connect database")
	}

	// Получаем экземпляр базы данных для проверки соединения
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Failed to get database instance: %v\n", err)
		panic("Failed to get database instance")
	}

	// Проверяем соединение с базой данных
	err = sqlDB.Ping()
	if err != nil {
		log.Printf("Failed to ping database: %v\n", err)
		panic("Failed to ping database")
	}

	// First, migrate without the name field constraint
	err = DB.AutoMigrate(
		&models.User{},    // Миграция модели пользователя
		&models.News{},    // Миграция модели новости
		&models.Comment{}, // Миграция модели комментария
	)
	if err != nil {
		log.Printf("Failed to migrate database: %v\n", err)
		panic("Failed to migrate database")
	}

	// Update existing users to have a default name if it's empty
	err = DB.Model(&models.User{}).Where("name = ? OR name IS NULL", "").Updates(map[string]interface{}{
		"name": "Anonymous User",
	}).Error
	if err != nil {
		log.Printf("Failed to update existing users: %v\n", err)
		panic("Failed to update existing users")
	}

	// Now add the NOT NULL constraint
	err = DB.Exec("ALTER TABLE users ALTER COLUMN name SET NOT NULL").Error
	if err != nil {
		log.Printf("Failed to add NOT NULL constraint to name column: %v\n", err)
		panic("Failed to add NOT NULL constraint to name column")
	}

	// Логируем успешное подключение
	log.Println("Successfully connected to database")
}
