// pkg/api/nextdate.go
package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

// afterNow возвращает true, если date > now (только дата, без времени)
func afterNow(date time.Time, now time.Time) bool {
	y1, m1, d1 := date.Date()
	y2, m2, d2 := now.Date()
	if y1 > y2 {
		return true
	}
	if y1 < y2 {
		return false
	}
	if m1 > m2 {
		return true
	}
	if m1 < m2 {
		return false
	}
	return d1 > d2
}

// NextDate вычисляет следующую дату по правилу повторения
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat is empty")
	}

	// Парсим начальную дату
	start, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", errors.New("invalid dstart format")
	}

	parts := strings.Split(repeat, " ")
	switch parts[0] {
	case "y":
		// Ежегодно
		if len(parts) != 1 {
			return "", errors.New("invalid 'y' format")
		}
		next := start
		for !afterNow(next, now) {
			next = next.AddDate(1, 0, 0)
		}
		return next.Format(dateFormat), nil

	case "d":
		// Повтор каждые N дней
		if len(parts) != 2 {
			return "", errors.New("invalid 'd' format: missing number")
		}
		days, err := strconv.Atoi(parts[1])
		if err != nil || days <= 0 || days > 400 {
			return "", errors.New("invalid day interval: must be 1-400")
		}
		next := start
		for !afterNow(next, now) {
			next = next.AddDate(0, 0, days)
		}
		return next.Format(dateFormat), nil

	default:
		// Пока не поддерживаем w и m
		return "", errors.New("unsupported repeat rule")
	}
}

// nextDateHandler — обработчик /api/nextdate
func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeatStr := r.FormValue("repeat")

	if dateStr == "" || repeatStr == "" {
		http.Error(w, "missing 'date' or 'repeat' parameter", http.StatusBadRequest)
		return
	}

	// Определяем now
	var now time.Time
	if nowStr == "" {
		now = time.Now()
	} else {
		parsed, err := time.Parse(dateFormat, nowStr)
		if err != nil {
			http.Error(w, "invalid 'now' format", http.StatusBadRequest)
			return
		}
		now = parsed
	}

	// Вычисляем следующую дату
	result, err := NextDate(now, dateStr, repeatStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}
