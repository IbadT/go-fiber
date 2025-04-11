package database

// import (
// 	"context"
// 	"log"

// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// var DB *pgxpool.Pool

// func InitDB(dsn string) {
// 	var err error

// 	DB, err = pgxpool.New(context.Background(), dsn)

// 	if err != nil {
// 		log.Fatal("Ошибка подключения к базе данных", err)
// 	}

// 	log.Println("Успешное подключение к PostgreSQL!")
// }

// func CloseDB() {
// 	DB.Close()
// 	log.Println("Соединение с PostgreSQL закрыто.")
// }
