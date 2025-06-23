package db

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
