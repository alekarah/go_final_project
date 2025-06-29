// Package api предоставляет HTTP обработчики для REST API планировщика задач.
//
// API поддерживает полный CRUD для задач и дополнительные операции:
//   - POST /api/task - добавление задачи
//   - GET /api/task?id=N - получение задачи по ID
//   - PUT /api/task - обновление задачи
//   - DELETE /api/task?id=N - удаление задачи
//   - GET /api/tasks - получение списка задач с поиском
//   - POST /api/task/done?id=N - отметка задачи как выполненной
//   - GET /api/nextdate - вспомогательный endpoint для тестирования
//
// Все endpoints возвращают JSON и поддерживают обработку ошибок.
// Поиск в /api/tasks поддерживает фильтрацию по тексту и дате в формате DD.MM.YYYY.
//
// Пример использования:
//
//	// Инициализация API маршрутов
//	api.Init()
//
//	// Запуск сервера
//	http.ListenAndServe(":7540", nil)
package api

import "net/http"

// Init регистрирует все API обработчики.
//
// Функция настраивает маршрутизацию для всех endpoints REST API.
// Должна быть вызвана до запуска HTTP сервера.
func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", doneTaskHandler)
}
