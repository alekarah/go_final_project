package api

import (
	"net/http"

	"todo/pkg/db"
)

// deleteTaskHandler обрабатывает DELETE запросы для удаления задач
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр id
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	// Удаляем задачу из базы данных
	err := db.DeleteTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	// Возвращаем пустой JSON в случае успеха
	writeJSON(w, map[string]any{})
}
