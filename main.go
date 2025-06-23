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
