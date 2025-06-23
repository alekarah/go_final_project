package nextdate

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// DateFormat - константа для формата дат
const DateFormat = "20060102"

// NextDate вычисляет следующую дату для повторяющейся задачи
// now - время, от которого ищется ближайшая дата
// dstart - исходное время в формате "20060102"
// repeat - правило повторения ("d <число>" или "y")
// Возвращает следующую дату в формате "20060102" и ошибку
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// Проверяем что правило повторения не пустое
	if repeat == "" {
		return "", errors.New("правило повторения не может быть пустым")
	}

	// Парсим исходную дату
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", errors.New("некорректная исходная дата: " + err.Error())
	}

	// Разбираем правило повторения
	parts := strings.Split(repeat, " ")

	switch parts[0] {
	case "y":
		// Ежегодное повторение
		if len(parts) != 1 {
			return "", errors.New("неверный формат правила для ежегодного повторения")
		}

		// Увеличиваем дату на год до тех пор, пока она не станет больше now
		for {
			date = date.AddDate(1, 0, 0)
			if AfterNow(date, now) {
				break
			}
		}

	case "d":
		// Повторение через указанное количество дней
		if len(parts) != 2 {
			return "", errors.New("неверный формат правила для повторения по дням")
		}

		// Парсим количество дней
		interval, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", errors.New("некорректное количество дней: " + err.Error())
		}

		// Проверяем ограничения
		if interval <= 0 || interval > 400 {
			return "", errors.New("количество дней должно быть от 1 до 400")
		}

		// Увеличиваем дату на указанное количество дней
		for {
			date = date.AddDate(0, 0, interval)
			if AfterNow(date, now) {
				break
			}
		}

	default:
		return "", errors.New("неизвестный тип правила повторения: " + parts[0])
	}

	// Возвращаем дату в формате "20060102"
	return date.Format(DateFormat), nil
}

// AfterNow проверяет, что первая дата больше второй (без учета времени)
func AfterNow(date, now time.Time) bool {
	// Сравниваем только даты, игнорируя время
	dateOnly := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	nowOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return dateOnly.After(nowOnly)
}
