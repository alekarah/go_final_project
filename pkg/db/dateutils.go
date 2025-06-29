package db

import (
	"fmt"
	"strings"
	"time"
)

// isDateFormat проверяет, соответствует ли строка формату DD.MM.YYYY.
//
// Валидация выполняется без регулярных выражений путем проверки
// длины строки, разделителей и символов.
func isDateFormat(s string) bool {
	// Проверяем длину строки (должна быть ровно 10 символов)
	if len(s) != 10 {
		return false
	}

	// Разделяем по точкам
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return false
	}

	// Проверяем длину каждой части
	if len(parts[0]) != 2 || len(parts[1]) != 2 || len(parts[2]) != 4 {
		return false
	}

	// Проверяем, что все части состоят только из цифр
	for _, part := range parts {
		for _, char := range part {
			if char < '0' || char > '9' {
				return false
			}
		}
	}

	return true
}

// convertDateFormat конвертирует дату из DD.MM.YYYY в YYYYMMDD.
//
// Используется для преобразования пользовательского ввода
// в формат базы данных.
func convertDateFormat(dateStr string) (string, error) {
	// Парсим дату в формате DD.MM.YYYY
	t, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		return "", fmt.Errorf("некорректный формат даты: %w", err)
	}

	// Возвращаем в формате YYYYMMDD
	return t.Format("20060102"), nil
}
