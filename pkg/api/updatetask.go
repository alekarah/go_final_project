package api

import (
	"encoding/json"
	"net/http"

	"todo/pkg/db"
)

// updateTaskHandler обрабатывает PUT запросы для обновления задач.
//
// Валидирует JSON данные (включая ID), проверяет обязательные поля
// и правила дат, обновляет задачу в БД.
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	// Десериализуем JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка десериализации JSON: " + err.Error()}, http.StatusBadRequest)
		return
	}

	// Проверяем обязательные поля
	if task.ID == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор задачи"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "Не указан заголовок задачи"}, http.StatusBadRequest)
		return
	}

	// Проверяем и корректируем дату (используем ту же логику что и при добавлении)
	err = checkDate(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	// Обновляем задачу в базе данных
	err = db.UpdateTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	// Возвращаем пустой JSON в случае успеха
	writeJSON(w, map[string]any{}, http.StatusOK)
}
