package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"todo/pkg/api"
)

const (
	defaultPort = "7540"
	webDir      = "./web"
)

// StartServer запускает веб-сервер
func StartServer() {
	// Получаем порт из переменной окружения или используем по умолчанию
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = defaultPort
	}

	// Инициализируем API обработчики
	api.Init()

	// Настраиваем файл-сервер для обслуживания статических файлов
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	// Запускаем сервер
	fmt.Printf("Сервер запущен на порту %s\n", port)
	fmt.Printf("Веб-интерфейс: http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
