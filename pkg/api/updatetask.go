package api

import (
	"encoding/json"
	"net/http"

	"todo/pkg/db"
)

// updateTaskHandler обрабатывает PUT запросы для обновления задач
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	// Десериализуем JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка десериализации JSON: " + err.Error()})
		return
	}

	// Проверяем обязательные поля
	if task.ID == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор задачи"})
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "Не указан заголовок задачи"})
		return
	}

	// Проверяем и корректируем дату (используем ту же логику что и при добавлении)
	err = checkDate(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	// Обновляем задачу в базе данных
	err = db.UpdateTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	// Возвращаем пустой JSON в случае успеха
	writeJSON(w, map[string]any{})
}
