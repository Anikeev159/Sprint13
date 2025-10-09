// pkg/api/gettask.go
package api

import (
	"go-final-project/pkg/db"
	"net/http"
	"strconv"
)

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "id is required"})
		return
	}

	// Проверяем, что id — число (опционально, но безопасно)
	if _, err := strconv.Atoi(id); err != nil {
		writeJSON(w, map[string]string{"error": "invalid id format"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": "task not found"})
		return
	}

	writeJSON(w, task)
}

//dasd
