// pkg/api/donetask.go
package api

import (
	"go-final-project/pkg/db"
	"net/http"
	"time"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "id is required"})
		return
	}

	// Получаем задачу
	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": "task not found"})
		return
	}

	// Если правило повторения пустое — удаляем
	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			writeJSON(w, map[string]string{"error": "failed to delete task"})
			return
		}
	} else {
		// Иначе — вычисляем следующую дату
		now := time.Now()
		nextDate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJSON(w, map[string]string{"error": "invalid repeat rule"})
			return
		}
		// Обновляем дату
		if err := db.UpdateDate(id, nextDate); err != nil {
			writeJSON(w, map[string]string{"error": "failed to update date"})
			return
		}
	}

	// Возвращаем пустой JSON {}
	writeJSON(w, map[string]interface{}{})
}
