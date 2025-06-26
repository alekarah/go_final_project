package db

// Tasks возвращает список задач, отсортированных по дате.
//
// Функция поддерживает поиск по тексту и дате:
//   - Пустой search - возвращает все задачи
//   - Поиск по дате в формате DD.MM.YYYY (например, "08.02.2024")
//   - Поиск по тексту в заголовке или комментарии задачи
//
// Параметры:
//   - limit: максимальное количество возвращаемых задач
//   - search: строка поиска (пустая строка для получения всех задач)
//
// Возвращает:
//   - []*Task: слайс указателей на задачи, отсортированный по дате
//   - error: ошибка выполнения запроса или nil при успехе
func Tasks(limit int, search string) ([]*Task, error) {
	var query string
	var args []any

	if search == "" {
		// Обычный запрос без поиска
		query = `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT ?`
		args = []any{limit}
	} else {
		// Проверяем, является ли search датой в формате DD.MM.YYYY
		if isDateFormat(search) {
			// Конвертируем дату из DD.MM.YYYY в YYYYMMDD
			dateStr, err := convertDateFormat(search)
			if err != nil {
				return nil, err
			}

			// Поиск по конкретной дате
			query = `SELECT id, date, title, comment, repeat FROM scheduler WHERE date = ? ORDER BY date ASC LIMIT ?`
			args = []any{dateStr, limit}
		} else {
			// Поиск по тексту в заголовке или комментарии
			searchPattern := "%" + search + "%"
			query = `SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date ASC LIMIT ?`
			args = []any{searchPattern, searchPattern, limit}
		}
	}

	// Выполняем запрос
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Создаем слайс для результата
	var tasks []*Task

	// Сканируем результаты
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	// Проверяем ошибки при итерации
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Если задач нет, возвращаем пустой слайс (не nil)
	if tasks == nil {
		tasks = []*Task{}
	}

	return tasks, nil
}
