// pkg/api/validation.go
package api

import (
	"errors"
	"os"
	"time"

	"go-final-project/pkg/db"
)

func getNow() time.Time {
	nowStr := os.Getenv("TODO_NOW")
	if nowStr != "" {
		if t, err := time.Parse(dateFormat, nowStr); err == nil {
			return t
		}
	}
	return time.Now()
}

func checkDate(task *db.Task) error {
	now := getNow()
	nowStr := now.Format(dateFormat)

	if task.Date == "" {
		task.Date = nowStr
	}

	t, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return errors.New("invalid date format")
	}

	if t.Format(dateFormat) < nowStr {
		if task.Repeat == "" {
			task.Date = nowStr
		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	}

	return nil
}
