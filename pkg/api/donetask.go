// pkg/api/donetask.go
package api

import (
	"net/http"
	"time"

	"go-final-project/pkg/db"
)

func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		WriteJSON(w, map[string]string{"error": "id is required"}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		WriteJSON(w, map[string]string{"error": "task not found"}, http.StatusNotFound)
		return
	}

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			WriteJSON(w, map[string]string{"error": "failed to delete task"}, http.StatusInternalServerError)
			return
		}
	} else {
		now := time.Now()
		nextDate, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			WriteJSON(w, map[string]string{"error": "invalid repeat rule"}, http.StatusBadRequest)
			return
		}

		if err := db.UpdateDate(id, nextDate); err != nil {
			WriteJSON(w, map[string]string{"error": "failed to update date"}, http.StatusInternalServerError)
			return
		}
	}

	WriteJSON(w, map[string]interface{}{}, http.StatusOK)
}
