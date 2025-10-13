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
		WriteJSON(w, map[string]string{"error": "id is required"}, http.StatusBadRequest)
		return
	}

	// Проверяем, что id — число (опционально, но безопасно)
	if _, err := strconv.Atoi(id); err != nil {
		WriteJSON(w, map[string]string{"error": "invalid id format"}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		WriteJSON(w, map[string]string{"error": "task not found"}, http.StatusNotFound)
		return
	}

	WriteJSON(w, task, http.StatusOK)
}

//dasd
