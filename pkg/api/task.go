package api

import "net/http"

// taskHandler обрабатывает запросы к /api/task с разными HTTP методами
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Добавление задачи
		addTaskHandler(w, r)
	case http.MethodGet:
		// Получение задачи по ID
		getTaskHandler(w, r)
	case http.MethodPut:
		// Обновление задачи
		updateTaskHandler(w, r)
	case http.MethodDelete:
		// Удаление задачи
		deleteTaskHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
