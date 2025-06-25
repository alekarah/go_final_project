package db

import "fmt"

// Task представляет задачу в системе
type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// AddTask добавляет задачу в базу данных и возвращает идентификатор
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

// GetTask получает задачу по идентификатору
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

// UpdateTask обновляет существующую задачу
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
