// pkg/db/db.go
package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite" // драйвер для SQLite
)

// Глобальная переменная для подключения к БД
var DB *sql.DB

// SQL-схема: таблица + индекс
const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT '',
    title VARCHAR(255) NOT NULL DEFAULT '',
    comment TEXT NOT NULL DEFAULT '',
    repeat VARCHAR(128) NOT NULL DEFAULT ''
);
CREATE INDEX idx_date ON scheduler(date);
`

func Init(dbFile string) error {
	log.Printf("Инициализация БД: %s", dbFile)

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Println("Файл БД не существует — создаём...")
		db, err := sql.Open("sqlite", dbFile)
		if err != nil {
			log.Printf("Ошибка открытия БД: %v", err)
			return err
		}
		defer db.Close()

		_, err = db.Exec(schema)
		if err != nil {
			log.Printf("Ошибка создания таблицы: %v", err)
			return err
		}
		log.Println("Таблица создана успешно")
		return nil
	}

	log.Println("Файл БД существует — открываем...")
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		log.Printf("Ошибка открытия существующей БД: %v", err)
		return err
	}

	// Проверяем соединение
	if err := db.Ping(); err != nil {
		log.Printf("Ошибка проверки соединения: %v", err)
		return err
	}

	DB = db
	log.Println("БД успешно инициализирована")
	return nil
}
