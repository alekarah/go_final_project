package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"todo/pkg/db"
)

// addTaskHandler обрабатывает POST запросы для добавления задач
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	// Десериализуем JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка десериализации JSON: " + err.Error()})
		return
	}

	// Проверяем обязательное поле title
	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "Не указан заголовок задачи"})
		return
	}

	// Проверяем и корректируем дату
	err = checkDate(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	// Добавляем задачу в базу данных
	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка добавления задачи: " + err.Error()})
		return
	}

	// Возвращаем ID созданной задачи
	writeJSON(w, map[string]string{"id": strconv.FormatInt(id, 10)})
}
