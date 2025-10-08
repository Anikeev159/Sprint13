// pkg/api/validation.go
package api

import (
	"errors"
	"go-final-project/pkg/db"
	"time"
)

func checkDate(task *db.Task) error {
	now := time.Now()
	nowStr := now.Format(dateFormat)

	// Если дата не указана — ставим сегодня
	if task.Date == "" {
		task.Date = nowStr
	}

	// Проверяем корректность формата
	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return errors.New("invalid date format")
	}

	// Если дата в прошлом
	if !AfterNow(t, now) {
		if task.Repeat == "" {
			// Без правила — ставим сегодня
			task.Date = nowStr
		} else {
			// С правилом — вычисляем следующую дату
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	}

	return nil
}
