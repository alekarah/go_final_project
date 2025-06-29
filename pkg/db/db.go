// Package db предоставляет функции для работы с базой данных SQLite.
//
// Пакет содержит структуры данных и операции для управления задачами
// в планировщике: добавление, получение, обновление и удаление.
//
// База данных использует SQLite с таблицей scheduler, которая содержит
// поля для хранения информации о задачах и правилах их повторения.
//
// Пример использования:
//
//	// Инициализация базы данных
//	err := db.Init("scheduler.db")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer db.Close()
//
//	// Добавление задачи
//	task := &db.Task{
//	    Title: "Важная задача",
//	    Date:  "20240201",
//	    Repeat: "d 7",
//	}
//	id, err := db.AddTask(task)
package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

// Глобальная переменная для хранения соединения с БД
var db *sql.DB

// schema содержит SQL команды для создания таблицы scheduler и индекса.
//
// Таблица включает поля:
//   - id: автоинкрементный первичный ключ
//   - date: дата задачи в формате YYYYMMDD
//   - title: заголовок задачи
//   - comment: комментарий к задаче
//   - repeat: правила повторения задачи
const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX idx_scheduler_date ON scheduler(date);
`

// Init инициализирует соединение с базой данных SQLite.
//
// Функция проверяет существование файла базы данных и создает
// таблицу scheduler с индексом, если файл не существует.
//
// Параметры:
//   - dbFile: путь к файлу базы данных SQLite
//
// Возвращает:
//   - error: ошибка инициализации или nil при успехе
func Init(dbFile string) error {
	// Проверяем существование файла БД
	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	// Открываем соединение с базой данных
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	// Проверяем соединение
	if err = db.Ping(); err != nil {
		return err
	}

	// Если файл БД не существовал, создаем таблицу и индекс
	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetDB возвращает активное соединение с базой данных.
//
// Возвращает:
//   - *sql.DB: указатель на соединение с базой данных
func GetDB() *sql.DB {
	return db
}

// Close закрывает соединение с базой данных.
//
// Возвращает:
//   - error: ошибка закрытия соединения или nil при успехе
func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
