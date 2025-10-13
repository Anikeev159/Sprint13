// pkg/api/tasks.go
package api

import (
	"encoding/json"
	"go-final-project/pkg/db"
	"net/http"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр search
	search := r.FormValue("search")

	// Ограничиваем 50 задачами
	limit := 50
	// Получаем задачи из БД
	tasks, err := db.Tasks(limit, search)
	if err != nil {
		writeJSON(w, map[string]string{"error": "database error"})
		return
	}

	writeJSON(w, TasksResp{Tasks: tasks})
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}
