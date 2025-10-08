// pkg/api/deletetask.go
package api

import (
	"go-final-project/pkg/db"
	"net/http"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "id is required"})
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeJSON(w, map[string]string{"error": "task not found"})
		return
	}

	writeJSON(w, map[string]interface{}{})
}
