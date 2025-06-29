// Планировщик задач - веб-сервер для управления TODO-листом.
//
// Приложение предоставляет REST API для работы с задачами:
// добавление, получение, изменение, удаление и отметка выполнения.
// Поддерживает повторяющиеся задачи с правилами повторения.
//
// Веб-интерфейс доступен по адресу http://localhost:7540
//
// Переменные окружения:
//
//	TODO_PORT   - порт веб-сервера (по умолчанию 7540)
//	TODO_DBFILE - путь к файлу базы данных (по умолчанию scheduler.db)
//
// Пример запуска:
//
//	go run main.go
//	TODO_PORT=8080 go run main.go
package main

import (
	"log"
	"os"

	"todo/pkg/db"
	"todo/pkg/server"
)

func main() {
	// Определяем путь к файлу базы данных
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		dbFile = "scheduler.db"
	}

	// Инициализируем базу данных
	err := db.Init(dbFile)
	if err != nil {
		log.Fatal("Ошибка инициализации базы данных:", err)
	}

	// Запускаем сервер
	server.StartServer()
}
