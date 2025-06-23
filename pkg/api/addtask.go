package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"todo/pkg/db"
	"todo/pkg/nextdate"
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

// checkDate проверяет и корректирует дату задачи
func checkDate(task *db.Task) error {
	now := time.Now()

	// Если дата не указана, используем сегодняшнюю
	if task.Date == "" {
		task.Date = now.Format(nextdate.DateFormat)
	}

	// Проверяем корректность формата даты
	t, err := time.Parse(nextdate.DateFormat, task.Date)
	if err != nil {
		return errors.New("дата представлена в формате, отличном от " + nextdate.DateFormat)
	}

	// Если указано правило повторения, проверяем его корректность
	var next string
	if task.Repeat != "" {
		next, err = nextdate.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return errors.New("правило повторения указано в неправильном формате: " + err.Error())
		}
	}

	// Если дата задачи в прошлом
	if nextdate.AfterNow(now, t) {
		if task.Repeat == "" {
			// Если правила повторения нет, берем сегодняшнюю дату
			task.Date = now.Format(nextdate.DateFormat)
		} else {
			// Иначе используем вычисленную следующую дату
			task.Date = next
		}
	}

	return nil
}

// writeJSON сериализует данные в JSON и отправляет ответ
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Ошибка сериализации JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
