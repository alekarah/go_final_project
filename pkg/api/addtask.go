package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"todo/pkg/db"
)

// addTaskHandler обрабатывает POST запросы для добавления задач.
//
// Валидирует JSON данные, проверяет обязательные поля и правила дат,
// добавляет задачу в БД и возвращает ID созданной записи.
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	// Десериализуем JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка десериализации JSON: " + err.Error()}, http.StatusBadRequest)
		return
	}

	// Проверяем обязательное поле title
	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "Не указан заголовок задачи"}, http.StatusBadRequest)
		return
	}

	// Проверяем и корректируем дату
	err = checkDate(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	// Добавляем задачу в базу данных
	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка добавления задачи: " + err.Error()}, http.StatusInternalServerError)
		return
	}

	// Возвращаем ID созданной задачи
	writeJSON(w, map[string]string{"id": strconv.FormatInt(id, 10)}, http.StatusCreated)
}
