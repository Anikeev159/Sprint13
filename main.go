// main.go
package main

import (
	"go-final-project/pkg/db"
	"go-final-project/server"
	"log"
	"os"
)

func main() {
	// Определяем путь к БД: из переменной окружения или по умолчанию
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	// Инициализируем БД
	if err := db.Init(dbFile); err != nil {
		log.Fatalf("Не удалось инициализировать БД: %v", err)
	}
	defer db.DB.Close()

	// Запускаем сервер
	server.Start()
}

//asd
