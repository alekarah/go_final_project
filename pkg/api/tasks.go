package api

import (
	"net/http"

	"todo/pkg/db"
)

// TasksResp структура ответа для списка задач
type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

// tasksHandler обрабатывает GET запросы к /api/tasks
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получаем параметр поиска
	search := r.FormValue("search")

	// Получаем список задач из базы данных (максимум 50)
	tasks, err := db.Tasks(50, search)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка получения задач: " + err.Error()})
		return
	}

	// Возвращаем список задач
	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}
