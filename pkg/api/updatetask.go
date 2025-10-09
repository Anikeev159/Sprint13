// pkg/api/updatetask.go
package api

import (
	"encoding/json"
	"go-final-project/pkg/db"
	"net/http"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{"error": "invalid JSON"})
		return
	}

	if task.ID == "0" {
		writeJSON(w, map[string]string{"error": "id is required"})
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "title is required"})
		return
	}

	// Проверка и коррекция даты (как в addTaskHandler)
	if err := checkDate(&task); err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeJSON(w, map[string]string{"error": "task not found"})
		return
	}

	// Возвращаем пустой JSON {}
	writeJSON(w, map[string]interface{}{})
}
