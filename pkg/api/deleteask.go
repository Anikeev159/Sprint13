// pkg/api/deletetask.go
package api

import (
	"net/http"

	"go-final-project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		WriteJSON(w, map[string]string{"error": "id is required"}, http.StatusBadRequest)
		return
	}

	if err := db.DeleteTask(id); err != nil {
		WriteJSON(w, map[string]string{"error": "task not found"}, http.StatusNotFound)
		return
	}

	WriteJSON(w, map[string]interface{}{}, http.StatusOK)
}
