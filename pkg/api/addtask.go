// pkg/api/addtask.go
package api

import (
	"encoding/json"
	"net/http"

	"go-final-project/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		WriteJSON(w, map[string]string{"error": "invalid JSON"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		WriteJSON(w, map[string]string{"error": "title is required"}, http.StatusBadRequest)
		return
	}

	if err := checkDate(&task); err != nil {
		WriteJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		WriteJSON(w, map[string]string{"error": "database error"}, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, map[string]int64{"id": id}, http.StatusCreated)
}
