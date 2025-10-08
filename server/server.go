// server/server.go
package server

import (
	"go-final-project/pkg/api" // ← импорт API
	"log"
	"net/http"
	"os"
)

const defaultPort = "7540"
const webDir = "./web"

func Start() {
	// Регистрируем API-обработчики
	api.Init()

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = defaultPort
	}

	// Обслуживаем статику из ./web
	fs := http.FileServer(http.Dir(webDir))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil)) // ← nil — это нормально!
}
