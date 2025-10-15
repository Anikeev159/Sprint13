// pkg/api/tasks.go
package api

import (
	"net/http"

	"go-final-project/pkg/db"
)

const maxTasks = 50

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	search := r.FormValue("search")
	limit := maxTasks

	tasks, err := db.Tasks(limit, search)
	if err != nil {
		WriteJSON(w, map[string]string{"error": "database error"}, http.StatusInternalServerError)
		return
	}

	WriteJSON(w, TasksResp{Tasks: tasks}, http.StatusOK)
}
