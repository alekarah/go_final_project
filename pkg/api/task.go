package api

import "net/http"

// taskHandler обрабатывает запросы к /api/task с разными HTTP методами
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Добавление задачи
		addTaskHandler(w, r)
	case http.MethodGet:
		// TODO: получение задачи (будет реализовано на следующих шагах)
		writeJSON(w, map[string]string{"error": "GET метод пока не реализован"})
	case http.MethodPut:
		// TODO: обновление задачи (будет реализовано на следующих шагах)
		writeJSON(w, map[string]string{"error": "PUT метод пока не реализован"})
	case http.MethodDelete:
		// TODO: удаление задачи (будет реализовано на следующих шагах)
		writeJSON(w, map[string]string{"error": "DELETE метод пока не реализован"})
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
