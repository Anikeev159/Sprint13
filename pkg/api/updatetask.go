// pkg/api/updatetask.go
package api

import (
	"encoding/json"
	"net/http"

	"go-final-project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		WriteJSON(w, map[string]string{"error": "invalid JSON"}, http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		WriteJSON(w, map[string]string{"error": "id is required"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		WriteJSON(w, map[string]string{"error": "title is required"}, http.StatusBadRequest)
		return
	}

	if err := checkDate(&task); err != nil { // ← без параметра now
		WriteJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		WriteJSON(w, map[string]string{"error": "failed to update task"}, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, map[string]interface{}{}, http.StatusOK)
}
