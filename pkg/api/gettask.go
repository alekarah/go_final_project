package api

import (
	"database/sql"
	"net/http"

	"todo/pkg/db"
)

// getTaskHandler обрабатывает GET запросы для получения задачи по ID.
//
// Получает ID из параметра запроса, находит задачу в БД
// и возвращает JSON с данными задачи.
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр id
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	// Получаем задачу из базы данных
	task, err := db.GetTask(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeJSON(w, map[string]string{"error": "Задача не найдена"})
		} else {
			writeJSON(w, map[string]string{"error": "Ошибка получения задачи: " + err.Error()})
		}
		return
	}

	// Возвращаем задачу в JSON формате
	writeJSON(w, task)
}
