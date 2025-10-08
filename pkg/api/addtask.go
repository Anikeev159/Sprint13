// pkg/api/addtask.go
package api

import (
	"encoding/json"
	"go-final-project/pkg/db"
	"net/http"
)

// addTaskHandler — обработчик POST /api/task
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	// Десериализация JSON
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{"error": "invalid JSON"})
		return
	}

	// Проверка обязательного поля
	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "title is required"})
		return
	}

	// Проверка и коррекция даты
	if err := checkDate(&task); err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	// Добавление в БД
	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "database error"})
		return
	}

	// Ответ: {"id": "123"}
	writeJSON(w, map[string]int64{"id": id})
}
