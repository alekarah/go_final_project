package api

import (
	"database/sql"
	"net/http"
	"time"

	"todo/pkg/db"
	"todo/pkg/nextdate"
)

// doneTaskHandler обрабатывает POST запросы для отметки задач как выполненных.
//
// Для одноразовых задач - удаляет из БД. Для повторяющихся задач -
// вычисляет следующую дату выполнения и обновляет запись.
func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// Проверяем правило повторения
	if task.Repeat == "" {
		// Одноразовая задача - удаляем
		err = db.DeleteTask(id)
		if err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка удаления задачи: " + err.Error()})
			return
		}
	} else {
		// Периодическая задача - вычисляем следующую дату
		now := time.Now()
		nextDate, err := nextdate.NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка вычисления следующей даты: " + err.Error()})
			return
		}

		// Обновляем дату в базе данных
		err = db.UpdateDate(nextDate, id)
		if err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка обновления даты: " + err.Error()})
			return
		}
	}

	// Возвращаем пустой JSON в случае успеха
	writeJSON(w, map[string]any{})
}
