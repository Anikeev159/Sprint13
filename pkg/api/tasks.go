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
<<<<<<< HEAD
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
=======
	// Получаем параметр search
	search := r.FormValue("search")

	// Ограничиваем 50 задачами
	limit := 50
	// Получаем задачи из БД
	tasks, err := db.Tasks(limit, search)
	if err != nil {
		writeJSON(w, map[string]string{"error": "database error"})
>>>>>>> 246564f55ef46378affc64b1c8dc75b4b0d31cd3
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
