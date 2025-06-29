package api

import (
	"net/http"
	"time"

	"todo/pkg/nextdate"
)

// nextDateHandler обрабатывает GET запросы к /api/nextdate.
//
// Вспомогательный endpoint для тестирования функции NextDate.
// Принимает параметры now, date, repeat и возвращает следующую дату.
func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получаем параметры запроса
	nowParam := r.FormValue("now")
	dateParam := r.FormValue("date")
	repeatParam := r.FormValue("repeat")

	// Проверяем обязательные параметры
	if dateParam == "" {
		http.Error(w, "Параметр 'date' обязателен", http.StatusBadRequest)
		return
	}
	if repeatParam == "" {
		http.Error(w, "Параметр 'repeat' обязателен", http.StatusBadRequest)
		return
	}

	// Определяем текущую дату
	var now time.Time
	if nowParam == "" {
		// Если now не указан, используем текущую дату
		now = time.Now()
	} else {
		// Парсим переданную дату
		var err error
		now, err = time.Parse(nextdate.DateFormat, nowParam)
		if err != nil {
			http.Error(w, "Некорректный формат параметра 'now': "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Вызываем функцию NextDate из пакета nextdate
	result, err := nextdate.NextDate(now, dateParam, repeatParam)
	if err != nil {
		http.Error(w, "Ошибка вычисления следующей даты: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Возвращаем результат
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
