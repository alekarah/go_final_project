package db

import "fmt"

// Task представляет задачу в системе планировщика.
//
// Поля:
//   - ID: уникальный идентификатор задачи
//   - Date: дата выполнения в формате YYYYMMDD
//   - Title: заголовок задачи (обязательное поле)
//   - Comment: комментарий к задаче
//   - Repeat: правило повторения ("y" для ежегодно, "d N" для каждые N дней)
type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// AddTask добавляет новую задачу в базу данных.
//
// Параметры:
//   - task: указатель на структуру Task с данными задачи
//
// Возвращает:
//   - int64: ID добавленной задачи
//   - error: ошибка, если операция не удалась
func AddTask(task *Task) (int64, error) {
	var id int64

	// SQL запрос для добавления задачи
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`

	// Выполняем запрос
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}

	// Получаем ID добавленной записи
	id, err = res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetTask получает задачу по идентификатору.
//
// Параметры:
//   - id: строковый идентификатор задачи
//
// Возвращает:
//   - *Task: указатель на структуру Task с данными задачи
//   - error: ошибка, если задача не найдена или произошла ошибка БД
func GetTask(id string) (*Task, error) {
	task := &Task{}

	// SQL запрос для получения задачи по ID
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`

	// Выполняем запрос и сканируем результат
	err := db.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateTask обновляет существующую задачу в базе данных.
//
// Все поля задачи будут обновлены значениями из переданной структуры.
//
// Параметры:
//   - task: указатель на структуру Task с новыми данными (должен содержать ID)
//
// Возвращает:
//   - error: ошибка, если задача не найдена или произошла ошибка БД
func UpdateTask(task *Task) error {
	// SQL запрос для обновления задачи
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`

	// Выполняем запрос
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return err
	}

	// Проверяем количество обновленных записей
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// Если ни одна запись не была обновлена, значит задача не найдена
	if count == 0 {
		return fmt.Errorf("задача с ID %s не найдена", task.ID)
	}

	return nil
}

// DeleteTask удаляет задачу по идентификатору.
//
// Параметры:
//   - id: строковый идентификатор задачи
//
// Возвращает:
//   - error: ошибка, если задача не найдена или произошла ошибка БД
func DeleteTask(id string) error {
	// SQL запрос для удаления задачи
	query := `DELETE FROM scheduler WHERE id = ?`

	// Выполняем запрос
	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	// Проверяем количество удаленных записей
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// Если ни одна запись не была удалена, значит задача не найдена
	if count == 0 {
		return fmt.Errorf("задача с ID %s не найдена", id)
	}

	return nil
}

// UpdateDate обновляет только дату задачи.
//
// Используется при отметке повторяющихся задач как выполненных
// для установки следующей даты выполнения.
//
// Параметры:
//   - nextDate: новая дата в формате YYYYMMDD
//   - id: строковый идентификатор задачи
//
// Возвращает:
//   - error: ошибка, если задача не найдена или произошла ошибка БД
func UpdateDate(nextDate string, id string) error {
	// SQL запрос для обновления только даты
	query := `UPDATE scheduler SET date = ? WHERE id = ?`

	// Выполняем запрос
	res, err := db.Exec(query, nextDate, id)
	if err != nil {
		return err
	}

	// Проверяем количество обновленных записей
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	// Если ни одна запись не была обновлена, значит задача не найдена
	if count == 0 {
		return fmt.Errorf("задача с ID %s не найдена", id)
	}

	return nil
}
